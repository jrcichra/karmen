use async_trait::async_trait;
use karmen::karmen_client::KarmenClient;
pub mod karmen {
    tonic::include_proto!("karmen");
}

struct Karmen {
    client: KarmenClient<tonic::transport::Channel>,
}
#[async_trait]
trait KarmenTraits {
    async fn new(host: &str, port: u16) -> Result<Karmen, tonic::transport::Error>;
    async fn ping(&mut self) -> Result<String, Box<dyn std::error::Error>>;
}

#[async_trait]
impl KarmenTraits for Karmen {
    async fn new(host: &str, port: u16) -> Result<Karmen, tonic::transport::Error> {
        let client = KarmenClient::connect(format!("http://{}:{}", host, port)).await?;
        Ok(Karmen { client })
    }
    async fn ping(&mut self) -> Result<String, Box<dyn std::error::Error>> {
        let request = tonic::Request::new(karmen::Ping {
            message: "Rusty Karmen!".into(),
        });
        let response = self.client.ping_pong(request).await?;
        Ok(response.into_inner().message)
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = Karmen::new("localhost", 8080).await?;
    let msg = client.ping().await?;
    println!("RESPONSE={:?}", msg);
    Ok(())
}
