use karmen::karmen_client::KarmenClient;
use std::collections::HashMap;
use std::sync::Arc;
use std::time::SystemTime;
use std::time::UNIX_EPOCH;
use tokio::sync::mpsc;
use tokio::sync::Mutex;

pub mod karmen {
    tonic::include_proto!("karmen");
}

struct Karmen {
    name: String,
    client: KarmenClient<tonic::transport::Channel>,
    actions: Arc<Mutex<HashMap<String, fn(parameters: HashMap<String, String>) -> karmen::Result>>>,
}

#[tonic::async_trait]
trait KarmenTraits {
    async fn new(name: &str, host: &str, port: u16)
        -> Result<Arc<Karmen>, tonic::transport::Error>;
    async fn ping(&self) -> Result<String, Box<dyn std::error::Error>>;
    async fn register(&self) -> Result<karmen::Result, Box<dyn std::error::Error>>;
    async fn handle_actions(&self) -> Result<(), Box<dyn std::error::Error>>;
    async fn add_action(
        &self,
        action: fn(parameters: HashMap<String, String>) -> karmen::Result,
        name: &str,
    ) -> Result<(), Box<dyn std::error::Error>>;
    async fn run_event(
        &self,
        name: &str,
        params: HashMap<String, String>,
    ) -> Result<karmen::Result, Box<dyn std::error::Error>>;
}

#[tonic::async_trait]
impl KarmenTraits for Karmen {
    async fn new(
        name: &str,
        host: &str,
        port: u16,
    ) -> Result<Arc<Karmen>, tonic::transport::Error> {
        let client = KarmenClient::connect(format!("http://{}:{}", host, port)).await?;
        Ok(Arc::new(Karmen {
            name: name.to_string(),
            client: client,
            actions: Arc::new(Mutex::new(HashMap::new())),
        }))
    }
    async fn ping(&self) -> Result<String, Box<dyn std::error::Error>> {
        let request = tonic::Request::new(karmen::Ping {
            message: "Rusty Karmen!".into(),
        });
        let response = self.client.clone().ping_pong(request).await?.into_inner();
        Ok(response.message)
    }
    async fn register(&self) -> Result<karmen::Result, Box<dyn std::error::Error>> {
        let request = tonic::Request::new(karmen::RegisterRequest {
            name: self.name.clone(),
            timestamp: SystemTime::now().duration_since(UNIX_EPOCH)?.as_secs() as i64,
            // no actions or events are sent as part of registration
            // TODO: investigate if these can be removed from the protobuf definition
            actions: HashMap::new(),
            events: HashMap::new(),
        });
        let response = self.client.clone().register(request).await?.into_inner();
        match response.result {
            Some(r) => {
                if r.code != 200 {
                    println!("{}", r.code);
                }
                Ok(r)
            }
            None => Err(Box::new(std::io::Error::new(
                std::io::ErrorKind::Other,
                "No response",
            ))),
        }
    }
    async fn handle_actions(&self) -> Result<(), Box<dyn std::error::Error>> {
        //send an action response with our details one time. For the rest of time, sleep
        let name = self.name.clone();
        let (tx, mut rx): (
            mpsc::Sender<karmen::ActionResponse>,
            mpsc::Receiver<karmen::ActionResponse>,
        ) = mpsc::channel(32);
        let request = async_stream::stream! {
            // Send our details once
            let response = karmen::ActionResponse {
                hostname: name.clone(),
                // unused on inital message
                request: None,
                result: None,
            };
            yield response;
            // Listen for more outgoing messages from the queue
            while let Some(response) = rx.recv().await {
                yield response;
            }
            println!("{}", "Closing stream");
        };
        let mut stream = self
            .client
            .clone()
            .action_dispatcher(tonic::Request::new(request))
            .await?
            .into_inner();
        while let Some(req) = stream.message().await? {
            println!("Got a request to run {:?}", req);
            // Prepare the function call with a clone of the request
            let action = req.clone().action.unwrap();
            let action_name = action.action_name.clone();
            let parameters = action.parameters;
            // Run the function
            let result = self.actions.lock().await.get(&action_name).unwrap()(parameters);
            // Process the result
            let response = karmen::ActionResponse {
                hostname: self.name.clone(),
                request: Some(req),
                result: Some(result),
            };
            // Send the result back through the queue
            tx.send(response).await?;
        }
        Ok(())
    }
    async fn add_action(
        &self,
        action: fn(parameters: HashMap<String, String>) -> karmen::Result,
        name: &str,
    ) -> Result<(), Box<dyn std::error::Error>> {
        self.actions.lock().await.insert(name.to_string(), action);
        Ok(())
    }
    async fn run_event(
        &self,
        name: &str,
        params: HashMap<String, String>,
    ) -> Result<karmen::Result, Box<dyn std::error::Error>> {
        println!("Running event {}", name);
        let event = Some(karmen::Event {
            event_name: name.to_string(),
            timestamp: SystemTime::now().duration_since(UNIX_EPOCH)?.as_secs() as i64,
            // TODO: determine if this is used
            parameters: HashMap::new(),
        });
        let request = tonic::Request::new(karmen::EventRequest {
            event: event,
            parameters: params,
            requester_name: self.name.clone(),
            // TODO: determine if this is used
            uuid: "".to_string(),
        });
        let response = self.client.clone().emit_event(request).await?.into_inner();
        match response.result {
            Some(r) => {
                if r.code != 200 {
                    println!("{}", r.code);
                }
                Ok(r)
            }
            None => Err(Box::new(std::io::Error::new(
                std::io::ErrorKind::Other,
                "No response",
            ))),
        }
    }
}

fn a_fn(parameters: HashMap<String, String>) -> karmen::Result {
    println!("Running action a_fn");
    //sleep
    let sleep_time = parameters.get("seconds").unwrap().parse::<u64>().unwrap();
    println!("Sleeping for {} seconds", sleep_time);
    std::thread::sleep(std::time::Duration::from_secs(sleep_time));
    karmen::Result {
        code: 200,
        parameters: parameters,
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = Karmen::new("bob", "localhost", 8080).await?;
    client.add_action(a_fn, "sleep").await?;
    client.register().await?;
    let client2 = client.clone();
    let join = tokio::spawn(async move {
        let _ = client2.handle_actions().await;
    });
    client.run_event("pleaseSleep", HashMap::new()).await?;
    join.await?;
    Ok(())
}
