use std::env;
use env_config::generate_documentation;
use tracing_subscriber::fmt;
use tracing::{info, Level};

fn main() {
    let subscriber = fmt::Subscriber::builder()
        .with_max_level(Level::TRACE)
        .with_timer(fmt::time::Uptime::default())
        .finish();

    tracing::subscriber::set_global_default(subscriber)
        .expect("setting tracing default failed");

    const FILENAME: &str = "env_variables.md";

    info!("Creating env variable documentation");

    generate_documentation(FILENAME).expect("Failed to generate documentation");

    info!("Documentation created at {}/{}", env::current_dir().unwrap().display() , FILENAME);
}
