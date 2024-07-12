use tracing_subscriber::fmt::format::FmtSpan;

pub fn init() {
    let evn_filter =
        std::env::var("RUST_LOG").unwrap_or("parcel_locker=info,warp=error".to_string());
    tracing_subscriber::fmt()
        .with_env_filter(evn_filter)
        .with_span_events(FmtSpan::CLOSE)
        .init();
}

pub fn construct_tracing_span_for_request(
) -> warp::trace::Trace<fn(warp::trace::Info) -> tracing::Span> {
    warp::trace(|info| {
        let mut remote_addr = String::new();
        if let Some(remote_socket_addr) = info.remote_addr() {
            remote_addr = format!("{remote_socket_addr}");
        }
        tracing::info_span!(
            "request",
            request_id = %uuid::Uuid::new_v4(),
            method = %info.method(),
            path = %info.path(),
            remote_addr = remote_addr,
            version = ?info.version(),
            referer = %info.referer().unwrap_or(""),
            user_agent = %info.user_agent().unwrap_or(""),
            host = %info.host().unwrap_or(""),
            request_headers = ?info.request_headers()
        )
    })
}
