pub fn error_message(s: &str) {
    println!("\n❌ {} has not been installed correctly\n", s);
}

// pub fn skip_message(s: &str) {
//     println!("\n⏩ {} installation has been skipped\n", s);
// }

pub fn success_message(s: &str) {
    println!("\n✅ {} has been installed\n", s);
}

// pub fn start_success_message(s: &str) {
//     println!("\n✅ {} has been started\n", s);
// }
//
// pub fn already_installed_message(s: &str) {
//     println!("\n✅ {} is already installed\n", s);
// }

pub fn installing_dependency(s: &str) {
    println!("\n🛠  {} is installing\n", s);
}

// pub fn updating_dependency(s: &str) {
//     println!("\n🛠  {} is updating\n", s);
// }
