use crate::domain::model::ParcelLocker;
use mockall::automock;

pub enum LoadError {
    NotFound,
    StorageError(redis::RedisError),
}

#[automock]
pub trait Loading {
    fn load_parcel_locker_by_id(&self, id: &str) -> Result<ParcelLocker, LoadError>;
    fn load_parcel_lockers(
        &self,
        page: usize,
        per_page: usize,
    ) -> Result<Vec<ParcelLocker>, LoadError>;
}
