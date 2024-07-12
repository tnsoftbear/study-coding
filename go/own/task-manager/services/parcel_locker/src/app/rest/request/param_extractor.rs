use std::collections::HashMap;

pub const PAGE_PARAM: &str = "page";
pub const PER_PAGE_PARAM: &str = "per_page";
const PAGE_DEF: usize = 1;
const PER_PAGE_DEF: usize = 10;

#[derive(Debug, PartialEq)]
pub struct PaginationParams {
    pub page: usize,
    pub per_page: usize,
}

#[derive(Debug, PartialEq)]
pub enum PaginationParamError {
    NotPositiveIntError((String, std::num::ParseIntError)),
    //    NotPositiveNumber(String),
}

pub fn extract_pagination(
    params: HashMap<String, String>,
) -> Result<PaginationParams, PaginationParamError> {
    let page = match params.get(PAGE_PARAM) {
        Some(page_param) => match page_param.parse::<usize>() {
            // Ok(page) if page > 0 => page,
            Ok(page) if page == 0 => PAGE_DEF,
            Ok(page) => page,
            Err(err) => {
                return Err(PaginationParamError::NotPositiveIntError((
                    PAGE_PARAM.to_string(),
                    err,
                )))
            }
        },
        None => PAGE_DEF,
    };
    let per_page = match params.get(PER_PAGE_PARAM) {
        Some(per_page_param) => match per_page_param.parse::<usize>() {
            // Ok(per_page) if per_page > 0 => per_page,
            // Ok(per_page) if per_page <= 0 => return Err(PaginationParamError::NotPositiveNumber("per_page".to_string())),
            Ok(per_page) if per_page == 0 => PER_PAGE_DEF,
            Ok(per_page) => per_page,
            Err(err) => {
                return Err(PaginationParamError::NotPositiveIntError((
                    PER_PAGE_PARAM.to_string(),
                    err,
                )))
            }
        },
        None => PER_PAGE_DEF,
    };
    Ok(PaginationParams { page, per_page })
}

#[cfg(test)]
mod param_extractor_test {
    use super::*;

    #[test]
    fn extract_pagination_default_values() {
        let params = HashMap::new();
        let result = extract_pagination(params);
        assert_eq!(
            result,
            Ok(PaginationParams {
                page: PAGE_DEF,
                per_page: PER_PAGE_DEF
            })
        );
    }

    #[test]
    fn extract_pagination_default_values_because_of_zeros() {
        let mut params = HashMap::new();
        params.insert(PAGE_PARAM.to_string(), "0".to_string());
        params.insert(PER_PAGE_PARAM.to_string(), "0".to_string());
        let result = extract_pagination(params);
        assert_eq!(
            result,
            Ok(PaginationParams {
                page: PAGE_DEF,
                per_page: PER_PAGE_DEF
            })
        );
    }

    #[test]
    fn extract_pagination_success() {
        let mut params = HashMap::new();
        params.insert(PAGE_PARAM.to_string(), "2".to_string());
        params.insert(PER_PAGE_PARAM.to_string(), "5".to_string());
        let result = extract_pagination(params);
        assert_eq!(
            result,
            Ok(PaginationParams {
                page: 2,
                per_page: 5
            })
        );
    }

    #[test]
    fn extract_pagination_invalid_page() {
        let mut params = HashMap::new();
        params.insert(PAGE_PARAM.to_string(), "-1".to_string());
        let result = extract_pagination(params);
        assert!(
            matches!(result, Err(PaginationParamError::NotPositiveIntError((key, _))) if key == PAGE_PARAM)
        );
    }

    #[test]
    fn extract_pagination_invalid_per_page() {
        let mut params = HashMap::new();
        params.insert(PER_PAGE_PARAM.to_string(), "-20".to_string());
        let result = extract_pagination(params);
        assert!(
            matches!(result, Err(PaginationParamError::NotPositiveIntError((key, _))) if key == PER_PAGE_PARAM)
        );
    }
}
