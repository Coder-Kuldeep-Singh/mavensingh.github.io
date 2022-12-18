// make div gloabally accessable
var containingItem = document.getElementById('digitalClock');


// give us the zero leading values
function ISODateString(d) {
    function pad(n) {
        return n < 10 ? '0' + n : n
    }
    return d.getFullYear() + '-' +
        pad(d.getMonth() + 1) + '-' +
        pad(d.getDate()) + ' ' +
        pad(d.getHours()) + ':' +
        pad(d.getMinutes()) + ':' +
        pad(d.getSeconds()) + ' ' +
        pad(d.getHours() >= 12 ? 'PM' : 'AM')
}

let doc = () => {
    let t = new Date();
    containingItem.innerHTML = ISODateString(t);
}

// set interval 1 sec so our clock
// our clock output can update on each
// second
setInterval(() => { doc() }, 1000);