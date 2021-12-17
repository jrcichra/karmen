use std::collections::HashMap;
use karmen::KarmenTraits;
use karmen::Karmen;

fn a_fn(parameters: HashMap<String, String>) -> karmen::karmen::Result {
    println!("Running action a_fn");
    //sleep
    let sleep_time = parameters.get("seconds").unwrap().parse::<u64>().unwrap();
    println!("Sleeping for {} seconds", sleep_time);
    std::thread::sleep(std::time::Duration::from_secs(sleep_time));
    karmen::karmen::Result{
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
