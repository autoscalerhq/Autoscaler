mod api_config;
mod worker_config;
use config::{Config };


// use crate::api_config::API_ENV_VARS;
// use crate::worker_config::WORKER_ENV_VARS;
mod shared_setup_config;

use serde::{Serialize, Deserialize};
use std::fs::{File, OpenOptions};
use std::io::Write;
// use tracing::{debug, trace, error, info, span, Level};
use std::error::Error;
use std::fmt;
use tracing::error;
use crate::shared_setup_config::SHARED_ENV_VARS;

#[derive(Debug, PartialEq, Eq, Hash, Clone, Serialize, Deserialize)]
pub enum EnvType {
    String,
    Boolean
}

#[derive(Debug, Hash, Serialize, Deserialize)]
pub enum EnvDefault {
    DefaultString(Option<&'static str>),
    DefaultBoolean(Option<bool>),
}

impl fmt::Display for EnvDefault {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            EnvDefault::DefaultString(Some(s)) => write!(f, "{}", s),
            EnvDefault::DefaultString(None) => write!(f, "None"),
            EnvDefault::DefaultBoolean(Some(b)) => write!(f, "{}", b),
            EnvDefault::DefaultBoolean(None) => write!(f, "None"),
        }
    }
}

#[derive(Debug, Hash, Serialize, Deserialize)]
pub struct EnvVar {
    pub section: &'static str,
    pub name: &'static str,
    pub description: &'static str,
    pub input_type: EnvType,
    pub default: EnvDefault
}

/// This function exports predefined environment variables to a JSON file.
///
/// # Arguments
///
/// * `env_vars` - A slice of environment variables to be written to the file.
/// * `service_name` - A string slice that holds the name of the service.
///   This will be used to determine the name of the file.
///
/// # Example
///
/// ```rust
/// use env_config::{EnvType, EnvVar, export_env_vars};
/// use env_config::EnvDefault::DefaultString;
/// let env_vars = vec![
///     EnvVar {
///         section: "API",
///         name: "PUBLIC".into(),
///         description: "Sets the api to be public",
///         input_type: EnvType::Boolean,
///         default: DefaultString(Some("true"))
///     },
///     EnvVar {
///         section: "API",
///         name: "ACTIVE".into(),
///         description: "Activates the api",
///         input_type: EnvType::Boolean,
///         default: DefaultString(Some("true"))},
///     EnvVar {
///         section: "Config",
///         name: "LOG_LEVEL".into(),
///         description: "Sets the logging level of the app",
///         input_type: EnvType::String,
///         default: DefaultString(None)},
/// ];
///
/// export_env_vars(&env_vars, "api").expect("panic message");
/// ```
///
/// This would create a `api_env_data.json` file with the environment variables.
///
/// # Errors
///
/// This function will return an error if the file cannot be created, or the data can't be written to the file.
pub fn export_env_vars(env_vars: &[EnvVar], service_name: &str) -> Result<(), Box<dyn Error>> {

    let filename = format!("{}_env_data.json", service_name);

    let file = match File::create(&filename) {
        Ok(file) => file,
        Err(e) => {
            error!("Failed to create file: {}", e);
            return Err(e.into());
        },
    };

    match serde_json::to_writer(file, env_vars) {
        Ok(_) => Ok(()),
        Err(e) => {
            error!("Failed to write to file: {}", e);
            Err(e.into())
        },
    }
}

pub fn generate_documentation(output_filename: &str) -> std::io::Result<()> {

    let mut file = OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true)
        .open(output_filename)?;

    writeln!(file, "# Environment Variables")?;

    let mut current_section = "";
    for var in SHARED_ENV_VARS {
        if var.section != current_section {
            writeln!(file, "\n## {}\n", var.section)?;
            current_section = var.section;
        }
        writeln!(
            file,
            "### {}: Type: {:?} (Default: {})-  {}",
            var.name, var.input_type, var.default, var.description,
        )?;
    }
    Ok(())
}


pub enum EnvironmentGroup {
    Api,
    Worker,
    Shared,
}

fn get_env_vars_by_group(group: EnvironmentGroup) -> &'static [EnvVar] {
    match group {
        // EnvironmentGroup::Api => &API_ENV_VARS,
        // EnvironmentGroup::Worker => &WORKER_ENV_VARS,
        EnvironmentGroup::Shared => &SHARED_ENV_VARS,
        _ => {&[]}
    }
}

pub fn get_env_var_value(group: EnvironmentGroup, name: &str, should_panic: bool) -> Option<String> {

    let builder = Config::builder().build().unwrap();

    if should_panic {
        panic!("Environment variable {} not found and no default provided!", name);
    }

    match builder.get_string(name){
        Ok(value) => Some(value),
        Err(_) => {
            // If environment variable was not found try to get a default value from the group
            for env_var in get_env_vars_by_group(group).iter() {
                if env_var.name == name {
                    // Assuming EnvVar has a field `default` of type Option<String>
                    return Some(env_var.default.to_string());
                }
            }

            // If no default was found either - panic if should_panic is true
            if should_panic {
                panic!("Environment variable {} not found and no default provided!", name);
            }

            None
        },
    }

}
