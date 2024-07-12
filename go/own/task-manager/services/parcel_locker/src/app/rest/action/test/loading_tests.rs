use crate::app::rest::action::loading::*;
use crate::app::rest::rejection::app_errors;
use crate::domain::model::ParcelLocker;
use crate::domain::repository::{LoadError, MockLoading};
use mockall::predicate::*;
use redis::RedisError;
use std::io;
use warp::http::{StatusCode, Version};
use warp::Reply;

#[tokio::test]
async fn load_parcel_locker_by_id_found() {
    // Arrange
    const EXISTING_ID: &str = "11111111";
    let mut loader = MockLoading::new();
    let expected_parcel_locker = ParcelLocker {
        id: EXISTING_ID.to_string(),
        name: "name-1".to_string(),
        latitude: 0.0,
        longitude: 0.0,
    };
    let expected_parcel_locker_clone = expected_parcel_locker.clone();
    loader
        .expect_load_parcel_locker_by_id()
        .with(eq(EXISTING_ID))
        .returning(move |_| Ok(expected_parcel_locker_clone.clone()));
    // Act
    let reply_result = load_parcel_locker_by_id(EXISTING_ID.to_string(), loader).await;
    // Assert
    assert!(reply_result.is_ok());
    let (parts, body) = reply_result.unwrap().into_response().into_parts();
    assert_eq!(parts.status, StatusCode::OK);
    assert_eq!(parts.version, Version::HTTP_11);
    assert!(parts.headers.contains_key("content-type"));
    assert_eq!(
        parts.headers.get("content-type").unwrap(),
        "application/json"
    );
    let body_bytes = hyper::body::to_bytes(body).await.unwrap();
    let actual_parcel_locker: ParcelLocker = serde_json::from_slice(&body_bytes).unwrap();
    assert_eq!(actual_parcel_locker, expected_parcel_locker);
}

#[tokio::test]
async fn load_parcel_locker_by_id_not_found() {
    // Arrange
    const ABSENT_ID: &str = "11111112";
    let mut loader = MockLoading::new();
    loader
        .expect_load_parcel_locker_by_id()
        .with(eq(ABSENT_ID))
        .returning(|_| Err(LoadError::NotFound));
    // Act
    let reply_result = load_parcel_locker_by_id(ABSENT_ID.to_string(), loader).await;
    // Assert
    assert!(reply_result.is_ok());
    let (parts, body) = reply_result.unwrap().into_response().into_parts();
    assert_eq!(parts.status, StatusCode::OK);
    assert_eq!(parts.version, Version::HTTP_11);
    assert!(parts.headers.contains_key("content-type"));
    assert_eq!(
        parts.headers.get("content-type").unwrap(),
        "application/json"
    );
    let actual_body_bytes = hyper::body::to_bytes(body).await.unwrap();
    assert_eq!(actual_body_bytes, hyper::body::Bytes::from("null"));
}

#[tokio::test]
async fn load_parcel_locker_by_id_failed_with_redis_error() {
    // Arrange
    const ID: &str = "11111112";
    const ERROR_MESSAGE: &str = "error message";
    let mut loader = MockLoading::new();
    loader
        .expect_load_parcel_locker_by_id()
        .with(eq(ID))
        .returning(|_| {
            Err(LoadError::StorageError(RedisError::from(io::Error::new(
                io::ErrorKind::Other,
                ERROR_MESSAGE.to_string(),
            ))))
        });
    // Act
    let reply_result = load_parcel_locker_by_id(ID.to_string(), loader).await;
    // Assert
    assert!(reply_result.is_err());
    if let Err(rejection) = reply_result {
        assert!(rejection.find::<app_errors::StorageError>().is_some());
        if let Some(err) = rejection.find::<app_errors::StorageError>() {
            assert_eq!(err.0.to_string(), ERROR_MESSAGE.to_string());
            println!("{:?}", err.0);
        } else {
            panic!("Expected StorageError");
        }
    } else {
        panic!("Expected Rejection");
    }
}
