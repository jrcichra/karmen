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
}

#[tonic::async_trait]
trait KarmenTraits {
    async fn new(name: &str, host: &str, port: u16) -> Result<Karmen, tonic::transport::Error>;
    async fn ping(&mut self) -> Result<String, Box<dyn std::error::Error>>;
    async fn register(&mut self) -> Result<karmen::Result, Box<dyn std::error::Error>>;
    async fn handle_actions(&mut self) -> Result<(), Box<dyn std::error::Error>>;
}

#[tonic::async_trait]
impl KarmenTraits for Karmen {
    async fn new(name: &str, host: &str, port: u16) -> Result<Karmen, tonic::transport::Error> {
        let client = KarmenClient::connect(format!("http://{}:{}", host, port)).await?;
        Ok(Karmen {
            name: name.to_string(),
            client,
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
            println!("{:?}", res);
        }
        Ok(())
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = Karmen::new("bob", "localhost", 8080).await?;
    client.register().await?;
    client.handle_actions().await?;
    Ok(())
}
