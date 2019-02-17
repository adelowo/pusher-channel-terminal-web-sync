(function() {
  const APP_KEY = '3b65aa197f334949f0ef';
  const APP_CLUSTER = 'eu';

  const logsDiv = document.getElementById('logs');

  const pusher = new Pusher(APP_KEY, {
    cluster: APP_CLUSTER,
  });

  const channel = pusher.subscribe('realtime-terminal');

  channel.bind('logs', data => {
    const divElement = document.createElement('div');
    divElement.innerHTML = data;

    logsDiv.appendChild(divElement);
  });
})();
