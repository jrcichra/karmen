use karmen::karmen_client::KarmenClient;
use karmen::Ping;

pub mod karmen {
    tonic::include_proto!("karmen");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = KarmenClient::connect("http://localhost:8080").await?;

    let request = tonic::Request::new(Ping {
        message: "Rusty Karmen!".into(),
    });

    let response = client.ping_pong(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}
