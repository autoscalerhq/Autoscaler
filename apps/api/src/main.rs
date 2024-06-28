mod main_test;

use actix_web::{get, middleware::Logger, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn index() -> impl Responder {
    println!("called / ");
    HttpResponse::Ok().body("Hello world!")
}

#[get("/again")]
async fn again() -> impl Responder {
    HttpResponse::Ok().body("Hello world again!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {

    println!("Starting HTTP server: go to http://localhost:8080");

    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(index)
            .service(again)
    })
        .bind(("0.0.0.0", 8080))?
        .run()
        .await
}

#[cfg(test)]
mod tests {
    use actix_web::{http::header::ContentType, test, App};

    use super::*;

    #[actix_web::test]
    async fn test_index_get() {
        let app = test::init_service(App::new().service(index)).await;
        let req = test::TestRequest::default()
            .insert_header(ContentType::plaintext())
            .to_request();
        let resp = test::call_service(&app, req).await;
        assert!(resp.status().is_success());
    }


}