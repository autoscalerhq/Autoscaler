use crate::{EnvType, EnvVar};
use crate::EnvDefault::DefaultString;

pub const SHARED_ENV_VARS: &[EnvVar] = &[
    EnvVar {
        section: "Pyroscope",
        name: "PYROSCOPE_URL",
        description: "The URL of the Pyroscope server.",
        input_type: EnvType::String,
        default: DefaultString(Some("http://localhost:4040" ) ),
    },
    EnvVar {
        section: "Pyroscope",
        name: "PYROSCOPE_USER",
        description: "The user to connect to the Pyroscope server.",
        input_type: EnvType::String,
        default: DefaultString(None ),
    },
    EnvVar {
        section: "Pyroscope",
        name: "PYROSCOPE_PASSWORD",
        description: "The pasword to authenticate to the Pyroscope server.",
        input_type: EnvType::String,
        default: DefaultString(None ),
    },
    EnvVar {
        section: "Pyroscope",
        name: "PYROSCOPE_SAMPLE_RAT",
        description: "The pasword to authenticate to the Pyroscope server.",
        input_type: EnvType::String,
        default: DefaultString(Some("100") ),
    },
    EnvVar {
        section: "Base",
        name: "APP_NAME",
        description: "The name of the application that is being instantiated.",
        input_type: EnvType::String,
        default: DefaultString(Some("App")),
    },


];

