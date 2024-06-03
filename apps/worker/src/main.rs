use env_config::{EnvironmentGroup, get_env_var_value};

#[shuttle_runtime::main]
async fn shuttle_main() -> Result<WorkerService, shuttle_runtime::Error> {
    println!("Start");
    Ok(WorkerService {})
}

// Customize this struct with things from `shuttle_main` needed in `bind`,
// such as secrets or database connections
struct WorkerService {}

#[shuttle_runtime::async_trait]
impl shuttle_runtime::Service for WorkerService {
    async fn bind(self, _addr: std::net::SocketAddr) -> Result<(), shuttle_runtime::Error> {
        println!("test");

        println!("test {:?}", get_env_var_value( EnvironmentGroup::Shared,"PYROSCOPE_URL", false));
        // Start your service and bind to the socket address
        Ok(())
    }
}