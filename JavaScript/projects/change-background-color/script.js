let doc = document.getElementById('submit');
doc.addEventListener('click', () => {
    let colorPicker = document.getElementById('userColor');
    document.body.style.backgroundColor = colorPicker.value;
});