var player;

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
  // Создание плеера
  player = new YT.Player("player", {
    events: {
      onReady: onPlayerReady,
      onStateChange: onPlayerStateChange,
    },
    videoId: "UONvpzG7yjo",
  });
}

function onPlayerReady(event) {
  document.getElementById("yt_status").innerHTML = "ready";
}

function onPlayerStateChange(event) {
  document.getElementById("yt_status").innerHTML = event.data;
  console.log(event);
}

// If the page is hidden, pause the video;
// if the page is shown, play the video
function handleVisibilityChange() {
  if (document[hidden]) {
    videoElement.pause();
    player.pauseVideo();
  } else {
    videoElement.play();
    player.playVideo();
  }
}
