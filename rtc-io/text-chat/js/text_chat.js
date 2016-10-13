// Set RTC options.
var rtcOpts = {
    room: 'test-data',
    signaller: 'http://10.0.2.246:3000',
    capture: false
};

// call RTC module
var rtc = RTC(rtcOpts);

// A contenteditable element to show our messages
var messages = document.getElementById('messages');

// Bind to events happening on the data channel
function bindDataChannelEvents(id, channel, attributes, connection) {

    // Receive message
    channel.onmessage = function (evt) {
        messages.innerHTML = evt.data;
    };

    // Send message
    messages.onkeyup = function () {
        channel.send(this.innerHTML);
    };
}

// Start working with the established session
function init(session) {
    session.createDataChannel('chat', {
        ordered: true,
        maxRetransmits: 12
    });
    session.on('channel:opened:chat', bindDataChannelEvents);
}

// Detect when RTC has established a session
rtc.on('ready', init);
