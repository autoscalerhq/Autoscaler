use pyroscope::pyroscope::PyroscopeAgentState;
use pyroscope::PyroscopeAgent;
use pyroscope_pprofrs::{pprof_backend, PprofConfig};

pub struct Profiling {
    pub agent: PyroscopeAgent<dyn PyroscopeAgentState>,
}

impl Profiling {

    pub fn new() -> Result<Profiling, Box<dyn std::error::Error>> {

        let pprof_config = PprofConfig::new().sample_rate(100);

        let backend_impl = pprof_backend(pprof_config);

        // Configure Pyroscope Agent
        let agent = PyroscopeAgent::builder(url, application_name.to_string())
            .basic_auth(user, password).backend(pprof_backend(PprofConfig::new().sample_rate(samplerate)))
            .tags([("app", "Rust"), ("TagB", "ValueB")].to_vec())
            .build()?;
    }
}