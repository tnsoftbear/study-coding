var ytPlayer;

// Set the name of the hidden property and the change event for visibility
var hidden, visibilityChange;
if (typeof document.hidden !== "undefined") {
  // Opera 12.10 and Firefox 18 and later support
  hidden = "hidden";
  visibilityChange = "visibilitychange";
} else if (typeof document.msHidden !== "undefined") {
  hidden = "msHidden";
  visibilityChange = "msvisibilitychange";
} else if (typeof document.webkitHidden !== "undefined") {
  hidden = "webkitHidden";
  visibilityChange = "webkitvisibilitychange";
}

function init() {
  initVideo();
}

function initVideo() {
  var videoElement = document.getElementById("videoElement");

  // Warn if the browser doesn't support addEventListener or the Page Visibility API
  if (
    typeof document.addEventListener === "undefined" ||
    hidden === undefined
  ) {
    console.log(
      "This demo requires a browser, such as Google Chrome or Firefox, that supports the Page Visibility API."
    );
  } else {
    // Handle page visibility change
    document.addEventListener(visibilityChange, handleVisibilityChange, false);

    // When the video pauses, set the title.
    // This shows the paused
    videoElement.addEventListener(
      "pause",
      function () {
        document.title = "Paused";
      },
      false
    );

    // When the video plays, set the title.
    videoElement.addEventListener(
      "play",
      function () {
        document.title = "Playing";
      },
      false
    );
  }
}

// Функция, вызываемая при загрузке API
// The API will call this function when the page has finished downloading the JavaScript for the player API
// https://developers.google.com/youtube/iframe_api_reference
function onYouTubeIframeAPIReady() {
  ytPlayer = new YT.Player("yt_player", {
    events: {
      onReady: () => {
        document.getElementById("yt_status").innerHTML = "ready";
      },
      onStateChange: onPlayerStateChange,
    },
    videoId: "UONvpzG7yjo",
  });
}

function onPlayerStateChange(event) {
  var statuses = {
    "-1": "unstarted",
    0: "ended",
    1: "playing",
    2: "paused",
    3: "buffering",
    5: "video cued",
  };

  document.getElementById("yt_status").innerHTML = statuses[event.data]
    ? statuses[event.data]
    : event.data;
}

// If the page is hidden, pause the video;
// if the page is shown, play the video
function handleVisibilityChange() {
  if (document[hidden]) {
    videoElement.pause();
    ytPlayer.pauseVideo();
  } else {
    videoElement.play();
    ytPlayer.playVideo();
  }
}
