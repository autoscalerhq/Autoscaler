use std::{env, str};
use std::os::unix::prelude::CommandExt;
use std::path::Path;
use std::process::{Command, exit};

use inquire::Select;

use crate::os_setup::set_up_os;
use crate::outputs::{error_message, installing_dependency, success_message};
use crate::tools::{
    check_and_install_git, install_docker, install_homebrew, install_nvm_mac, install_pnpm,
    start_docker_infra,
};

pub(crate) fn handle_installation_cmd(output: &Result<std::process::Output, std::io::Error>) {
    match output {
        Ok(_output) => {
            // git installation succeeded
            println!("Git installed successfully");
        }
        Err(e) => {
            // git installation failed
            println!("Failed to install Git: {}", e);
        }
    }
}

// Main function
pub fn setup() {
    const PLATFORM: &str = env::consts::OS;

    println!("Detecting Platform is: {}\n", PLATFORM);

    println!("Setting up defaults\n");
    set_up_os(PLATFORM);

    println!("\nInstalling Required Tools for {}\n", PLATFORM);

    check_hard_dependencies(PLATFORM);
    check_and_install_git(PLATFORM);

    match PLATFORM {
        "macos" => {
            install_homebrew();
            install_nvm_mac();
            install_pnpm();
            install_docker();
        }
        "windows" => {}
        _ => {}
    }

    start_docker_infra();
    refresh_shell(PLATFORM);
}

fn check_hard_dependencies(platform: &str) {
    match platform {
        "macOS" => {
            let output = Command::new("brew")
                .arg("--version")
                .output()
                .unwrap_or_else(|_| panic!("Failed to run Homebrew check"));

            let brew_check = str::from_utf8(&output.stdout)
                .expect("Could not decode byte array as UTF-8 string");

            if brew_check.is_empty() || output.status.code() != Some(0) {
                error_message("Homebrew");
                println!("⛔️ Homebrew is a hard dependency for this tool");

                installing_dependency("Homebrew");
                let output = Command::new("/bin/bash")
                    .arg("-c")
                    .arg(r#""$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)""#)
                    .output()
                    .expect("Failed to execute command");

                println!("{}", String::from_utf8_lossy(&output.stdout));
                success_message("Homebrew")
            }

            check_zsh();
        }
        "windows" => {}
        _ => {}
    }
}

fn check_zsh() {
    // Check if zsh is installed
    let output = Command::new("zsh")
        .arg("--version")
        .output()
        .unwrap_or_else(|_| panic!("Failed to check zsh"));

    // If zsh is not installed, don't go further.
    if output.status.code() != Some(0) {
        return;
    }

    // Zsh is installed, ask to install Oh My Zsh
    let user_choice = Select::new("Do you want to install Oh My Zsh?", vec!["Yes", "No"]).prompt();

    let response = user_choice.unwrap();

    // Check if user wants to install Oh My Zsh
    if response == "Yes" {
        let ohmyzsh_dir = format!("{}/.oh-my-zsh", env::var("HOME").unwrap());

        // Check if directory .oh-my-zsh exists in home directory
        if !Path::new(&ohmyzsh_dir).exists() {
            // Download and install Oh My Zsh!
            Command::new("sh")
                .arg("-c")
                .arg("$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)")
                .status()
                .expect("Failed to install Oh My Zsh!");

            // Check if Oh My Zsh! is installed successfully
            if !Path::new(&ohmyzsh_dir).exists() {
                println!("Error: Oh My Zsh!");
            } else {
                //success_message("Oh My Zsh!");
            }
        } else {
            println!("Oh My Zsh! is already installed.");
        }
    }
}

fn refresh_shell(platform: &str) {
    match platform {
        "macOS" => {
            let zprofile = env::var("ZPROFILE").expect("ZPROFILE is not set");

            let output = Command::new("sh")
                .arg("-c")
                .arg(format!("source {}", zprofile))
                .output()
                .expect("Failed to execute shell command");

            if output.status.success() {
                let shell = env::var("SHELL").expect("SHELL is not set");
                Command::new(shell).exec();
            } else {
                eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
                exit(1);
            }
        }
        "windows" => {
            let shell_profile = env::var("PROFILE").expect("PROFILE is not set");

            let output = Command::new("powershell")
                .arg("-Command")
                .arg(format!(". {}", shell_profile))
                .output()
                .expect("Failed to execute PowerShell command");

            if output.status.success() {
                Command::new("powershell").exec();
            } else {
                eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
                exit(1);
            }
        }
        _ => {
            println!("Unsupported sh")
        }
    }
}
