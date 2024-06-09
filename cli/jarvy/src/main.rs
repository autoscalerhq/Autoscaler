use clap::Parser;
use inquire::{InquireError, Select};

use crate::setup::setup;

mod os_setup;
mod outputs;
mod setup;
mod tools;

#[derive(Parser)]
struct Opts {
    #[clap(short, long)]
    pub run: bool,
}

fn main() {
    print_logo();

    println!(
        "\t\tHi, I'm Jarvy by Autoscaler!
Welcome to the codebase of the open-source autoscaling infrastructure for all!"
    );

    user_select();
}

fn user_select() {
    let options = vec![
        "Run the project",
        "Test the project",
        "Development environment setup",
    ];

    let selection: Result<&str, InquireError> =
        Select::new("What would you like to do today?", options).prompt();

    match selection {
        Ok(choice) => {
            println!("selection: {}", choice);
            match choice {
                "Run the project" => {
                    println!("R");
                    match std::process::Command::new("cargo").arg("run").output() {
                        Ok(output) => {
                            // Handle the output here
                            println!("Output: {}", String::from_utf8_lossy(&output.stdout));
                        }
                        Err(e) => println!("Failed to execute command: {}", e),
                    }
                }
                "Test the project" => {
                    println!("T");

                    match std::process::Command::new("cargo").arg("test").output() {
                        Ok(output) => {
                            // Handle the output here
                            println!("Output: {}", String::from_utf8_lossy(&output.stdout));
                        }
                        Err(e) => println!("Failed to execute command: {}", e),
                    }
                }
                "Development environment setup" => {
                    println!("D");
                    setup();
                }
                _ => {}
            }
        }
        Err(_) => {
            println!("No choice was made")
        }
    }
}

fn print_logo() {
    println!(
        "

  @@@                        @@@                        @@@
  @@@@@                     @@@@@                     @@@@@@
 @@@@@@@@                  @@@*@@@                  @@@@@@@@
  @@@@@@@@@@             @@@%-:-@@@              @@@@@@@@@@
   @@@@@%%@@@@          @@@%-::::%@@@          @@@@%%@@@@@
    @@@@@%=#@@@@       @@@%-:::::-%@@@       @@@@#=#@@@@@
     @@@@@%--#@@@@    @@@#::::::::-%@@@    @@@@*--%@@@@@
      @@@@@%---*@@@@ @@@*:::::::::::#@@@ @@@@+--=@@@@@@
       @@@@@@=---+@@@@@+:::::=#=:::::*@@@@@+---=@@@@@@
        @@@@@@=----#@@+:::::+@@@=:::::+@@*----=@@@@@
          @@@@@=--#@@=:::::=@@@@@=:::::+@@*--+@@@@@
           @@@@@+#@@=:::::-@@#:#@@=:::::=@@#*@@@@@
            @@@@@@@-:::::-@@#---%@%-:::::-@@@@@@@
             @@@@%-::::::%@%=----%@@-:::::-@@@@@
             @@@#---::--%@%=======@@%-::::--%@@@
            @@@#------=@@@+=======*@@@=------#@@@
           @@@#-----+@@@%@@@#+++#@@@%@@@+-----#@@@
          @@@*----*@@@@-::+@@@@@@@+::=@@@@*----*@@@
         @@@+---#@@@@@@@=:::=@@@=:::=@@@@@@@*---*@@@
        @@@+--#@@@@@@@@@@=:::::::::+@@@@@@@@@@#--+@@@
       @@@==%@@@@@@@@@@@@@+:::::::+@@@@@@@@@@@@@%=+@@@
     @@@%+@@@@@@@@    @@@@@*:::::#@@@@@    @@@@@@@%*@@@
    @@@@@@@@@@@@       @@@@@*:::#@@@@@       @@@@@@@@@@@@
   @@@@@@@@@@@          @@@@@#-#@@@@@          @@@@@@@@@@@
  @@@@@@@@@@             @@@@@@@@@@@             @@@@@@@@@@
  @@@@@@@@                @@@@@@@@@                @@@@@@@@@
 @@@@@@                    @@@@@@@                   @@@@@@@
  @@@                        @@@                        @@@
    "
    );
}
