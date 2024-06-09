use std::env;
use std::process::Command;
use std::str;

pub fn check_and_install_git(platform: &str) {
    match platform {
        "macos" => {
            let install_cmd = std::process::Command::new("brew")
                .arg("install")
                .arg("git")
                .output();
            crate::setup::handle_installation_cmd(&install_cmd);
        }
        "Linux" => {
            let install_cmd = std::process::Command::new("sudo")
                .arg("apt-get")
                .arg("install")
                .arg("-y")
                .arg("git")
                .output();
            crate::setup::handle_installation_cmd(&install_cmd);
        }
        "Windows" => {
            let install_cmd = std::process::Command::new("winget")
                .arg("install")
                .arg("Git")
                .output();
            crate::setup::handle_installation_cmd(&install_cmd);
        }
        _ => println!("Unsupported platform"),
    }
}

pub fn install_homebrew() {
    let test_brew_cmd = Command::new("brew")
        .arg("--version")
        .output()
        .expect("Failed to run brew");

    if !test_brew_cmd.status.success() {
        println!("Installing Homebrew");

        let curl_cmd = r#"/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)""#;
        let _output = Command::new("sh")
            .arg("-c")
            .arg(curl_cmd)
            .output()
            .expect("Failed to execute command");

        let zprofile = format!("{}/.zprofile", env::var("HOME").unwrap());
        let apple_chip_brew_bin = "/opt/homebrew/bin";
        let brew_bin = "/usr/local/bin";
        let entry = format!(
            "export PATH={}:{}:${}",
            brew_bin,
            apple_chip_brew_bin,
            env::var("PATH").unwrap()
        );

        let param_to_cmd = format!("grep -R {} {}", entry, zprofile);

        let _cmd = Command::new("sh")
            .arg("-c")
            .arg(param_to_cmd)
            .output()
            .expect("Failed to execute command");

        // if !cmd.status.success() {
        //     let mut file = OpenOptions::new()
        //         .append(true)
        //         .open(zprofile)
        //         .expect("Failed to open file");
        //
        //     //writeln!(file, "{}", entry).expect("Failed to write to file");
        // }

        let after_install_test_cmd = Command::new("brew")
            .arg("--version")
            .output()
            .expect("Failed to execute command");

        if !after_install_test_cmd.status.success() {
            eprintln!("Error: Homebrew");
        } else {
            println!("Successfully installed Homebrew");
        }
    } else {
        println!("Homebrew is already installed");
    }
}

pub fn install_nvm_mac() {
    match Command::new("nvm").arg("--version").output() {
        Ok(output) => {
            if output.status.success() {
                println!("NVM is already installed!");
            }
        }
        Err(_) => {
            // Attempt to install using brew
            let output = Command::new("brew")
                .arg("install")
                .arg("nvm")
                .output()
                .unwrap_or_else(|_| panic!("Failed to execute command"));

            // Check if the install command was successful
            if output.status.success() {
                println!("Successfully installed NVM!");
            } else {
                println!("Failed to install NVM...");
                if let Ok(output_str) = str::from_utf8(&output.stderr) {
                    println!("Error: {}", output_str);
                }
            }
        }
    }
}

pub fn install_pnpm() {
    let pnpm_version = "9.2.0";

    let test_pnpm_output = Command::new("pnpm")
        .arg("--version")
        .output()
        .expect("Failed to execute command");

    // If pnpm not found or any other problem occurred
    if !test_pnpm_output.status.success() {
        println!("Installing PNPM {}", pnpm_version);

        let npm_install_output = Command::new("npm")
            .arg("install")
            .arg("-g")
            .arg(format!("pnpm@{}", pnpm_version))
            .output()
            .expect("Failed to execute command");

        // Check if npm install command was successful
        if npm_install_output.status.success() {
            let after_install_test_output = Command::new("pnpm")
                .arg("--version")
                .output()
                .expect("Failed to execute command");

            // If pnpm command now runs properly
            if after_install_test_output.status.success() {
                println!("Successfully installed PNPM {}", pnpm_version);
            } else {
                println!("Error: PNPM");
            }
        }
    } else {
        println!("PNPM {} is already installed", pnpm_version);
    }
}

pub fn install_docker() {
    let check_homebrew_output = Command::new("brew")
        .arg("--version")
        .output()
        .expect("Failed to execute command");

    // If brew not found or any other problem occurred
    if check_homebrew_output.status.success() {
        let test_docker_output = Command::new("docker")
            .arg("--version")
            .output()
            .expect("Failed to execute command");

        // If docker not found or any other problem occurred
        if !test_docker_output.status.success() {
            println!("Installing Docker");
            let brew_install_output = Command::new("brew")
                .arg("install")
                .arg("docker")
                .output()
                .expect("Failed to execute command");

            // If Docker installed successfully
            if brew_install_output.status.success() {
                // Test Docker installation
                let after_install_test_output = Command::new("docker")
                    .arg("--version")
                    .output()
                    .expect("Failed to execute command");

                // If Docker now runs properly
                if after_install_test_output.status.success() {
                    println!("Successfully installed Docker");
                } else {
                    println!("Error: Docker");
                }
            }
        } else {
            println!("Docker is already installed");
        }
    } else {
        println!("Skipping Docker installation as Homebrew is not found");
    }
}

pub fn start_docker_infra() {
    let docker_compose_output = Command::new("docker-compose")
        .arg("-f")
        .arg("./docker//docker-compose.yml")
        .arg("up")
        .arg("-d")
        .output()
        .expect("Failed to execute command");

    if docker_compose_output.status.success() {
        println!("Successfully started Docker Infrastructure");
    } else {
        eprintln!(
            "An error occurred: \n\t {}. \nPlease run this from the root of your repository.",
            String::from_utf8_lossy(&docker_compose_output.stderr)
        );
    }
}
