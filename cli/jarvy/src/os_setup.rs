use std::process::{Command, Output};
use std::str;

use inquire::Select;

pub(crate) fn handle_output(output: &Output) {
    if !output.status.success() {
        eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
    }
}

fn get_cpu() -> String {
    let output = Command::new("uname")
        .arg("-m")
        .output()
        .expect("Failed to fetch CPU info");

    if output.status.success() {
        let s =
            str::from_utf8(&output.stdout).expect("Could not decode byte array as UTF-8 string");
        s.trim().to_string()
    } else {
        eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
        String::new()
    }
}

fn set_macosx_generics() {
    println!("Set MacOSx system configurations");

    let output = std::process::Command::new("defaults")
        .arg("write")
        .arg("com.apple.finder")
        .arg("AppleShowAllFiles")
        .arg("YES")
        .output()
        .expect("Failed to execute command");

    if !output.status.success() {
        eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
    }
}

pub fn set_up_os(platform: &str) {
    match platform {
        "macos" => {
            println!("Setting macos to show all file types.");

            let output = std::process::Command::new("defaults")
                .arg("write")
                .arg("com.apple.finder")
                .arg("AppleShowAllFiles")
                .arg("YES")
                .output()
                .expect("Failed to execute command");

            if !output.status.success() {
                eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
            }

            if get_cpu() != "arm64" {
                println!("Setting up Rosetta for arm64 emulation");

                let output = Command::new("softwareupdate")
                    .args(["--install-rosetta"])
                    .output()
                    .expect("Failed to start Rosetta installation");

                handle_output(&output);
            }

            set_macosx_generics();

            let xcode_check = Command::new("/usr/bin/xcode-select")
                .args(["-p"])
                .output()
                .expect("Failed to execute command");

            if xcode_check.status.success() {
                let update_xcode_prompt = Select::new(
                    "Xcode is already installed, do you want to update it?",
                    vec!["Yes", "No"],
                )
                .prompt();

                println!("\n");

                match update_xcode_prompt {
                    Ok(answer) => match answer {
                        "Yes" => {
                            let mut update = Command::new("softwareupdate")
                                .arg("-ia")
                                .spawn()
                                .expect("Failed to update software");

                            update.wait().expect("Failed to wait on child process");
                        }
                        "No" => {
                            println!("Xcode will not be updated.");
                        }
                        _ => unreachable!(),
                    },
                    Err(_) => {
                        println!("Could not read your response.");
                    }
                }
            } else {
                println!("Installing Xcode...");
                let mut install = Command::new("xcode-select")
                    .args(["--install"])
                    .spawn()
                    .expect("Failed to install Xcode");

                install.wait().expect("Failed to wait on child process");
            }
        }
        "Linux" => {
            println!("Nothing to configure");
        }
        "Windows" => {
            println!("Set Windows system configurations");

            let output = std::process::Command::new("powershell")
                .arg("/c")
                .arg("Set-ItemProperty -Path 'HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced' -Name 'Hidden' -Value 1")
                .output()
                .expect("Failed to execute command");

            if !output.status.success() {
                eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
            }
        }
        _ => println!("Unsupported platform"),
    }
    println!("\n");
}
