// use futures::stream::iter;

use karmen::karmen_client::KarmenClient;
use std::collections::HashMap;
use std::time::SystemTime;
use std::time::UNIX_EPOCH;

pub mod karmen {
    tonic::include_proto!("karmen");
}

struct Karmen {
    name: String,
    client: KarmenClient<tonic::transport::Channel>,
    actions: HashMap<String, fn(parameters: HashMap<String, String>) -> karmen::Result>,
}

#[tonic::async_trait]
trait KarmenTraits {
    async fn new(name: &str, host: &str, port: u16) -> Result<Karmen, tonic::transport::Error>;
    async fn ping(&mut self) -> Result<String, Box<dyn std::error::Error>>;
    async fn register(&mut self) -> Result<karmen::Result, Box<dyn std::error::Error>>;
    async fn handle_actions(&mut self) -> Result<(), Box<dyn std::error::Error>>;
    async fn add_action(
        &mut self,
        action: fn(parameters: HashMap<String, String>) -> karmen::Result,
        name: &str,
    ) -> Result<(), Box<dyn std::error::Error>>;
    async fn run_event(
        &mut self,
        name: &str,
        params: HashMap<String, String>,
    ) -> Result<karmen::Result, Box<dyn std::error::Error>>;
}

#[tonic::async_trait]
impl KarmenTraits for Karmen {
    async fn new(name: &str, host: &str, port: u16) -> Result<Karmen, tonic::transport::Error> {
        let client = KarmenClient::connect(format!("http://{}:{}", host, port)).await?;
        Ok(Karmen {
            name: name.to_string(),
            client,
            actions: HashMap::new(),
        })
    }
    async fn ping(&mut self) -> Result<String, Box<dyn std::error::Error>> {
        let request = tonic::Request::new(karmen::Ping {
            message: "Rusty Karmen!".into(),
        });
        let response = self.client.ping_pong(request).await?.into_inner();
        Ok(response.message)
    }
    async fn register(&mut self) -> Result<karmen::Result, Box<dyn std::error::Error>> {
        let request = tonic::Request::new(karmen::RegisterRequest {
            name: self.name.clone(),
            timestamp: SystemTime::now().duration_since(UNIX_EPOCH)?.as_secs() as i64,
            // no actions or events are sent as part of registration
            // TODO: investigate if these can be removed from the protobuf definition
            actions: HashMap::new(),
            events: HashMap::new(),
        });
        let response = self.client.register(request).await?.into_inner();
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
    async fn handle_actions(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        //empty request
        let request = async_stream::stream! {
            //TODO come up with a better solution for an empty stream
            while true {
                tokio::time::sleep(std::time::Duration::from_secs(12345)).await;
            }
            // never executed by design
            yield karmen::ActionResponse {
                request: None,
                result: None,
                hostname: "".to_string(),
            };
        };
        let mut stream = self.client.action_dispatcher(request).await?.into_inner();
        while let Some(res) = stream.message().await? {
            println!("Got a request to run {:?}", res);
        }
        Ok(())
    }
    async fn add_action(
        &mut self,
        action: fn(parameters: HashMap<String, String>) -> karmen::Result,
        name: &str,
    ) -> Result<(), Box<dyn std::error::Error>> {
        println!("Adding action {}", name);
        self.actions.insert(name.to_string(), action);
        Ok(())
    }
    async fn run_event(
        &mut self,
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
        let response = self.client.emit_event(request).await?.into_inner();
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
    println!("I am running sleep from rust!!!");
    karmen::Result {
        code: 200,
        parameters: parameters,
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = Karmen::new("bob", "localhost", 8080).await?;
    client.add_action(a_fn, "sleep").await?;
    client.register().await?;
    client.run_event("pleaseSleep", HashMap::new()).await?;
    client.handle_actions().await?;
    Ok(())
}
