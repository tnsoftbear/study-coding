use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Clone, Deserialize, Debug, Serialize, PartialEq)]
pub struct ParcelLocker {
    pub id: String,
    pub name: String,
    pub longitude: f64,
    pub latitude: f64,
}

impl ParcelLocker {
    pub fn to_tuples(&self) -> Vec<(&str, String)> {
        vec![
            ("id", self.id.clone()),
            ("name", self.name.clone()),
            ("longitude", self.longitude.to_string()),
            ("latitude", self.latitude.to_string()),
        ]
    }
}

impl From<HashMap<String, String>> for ParcelLocker {
    fn from(value: HashMap<String, String>) -> Self {
        ParcelLocker {
            id: value["id"].clone(),
            name: value["name"].clone(),
            longitude: value["longitude"].parse::<f64>().unwrap(),
            latitude: value["latitude"].parse::<f64>().unwrap(),
        }
    }
}
