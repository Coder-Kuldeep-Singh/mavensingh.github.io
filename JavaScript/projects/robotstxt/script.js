function addPath(type) {
    const pathGroup = document.createElement('div');
    pathGroup.classList.add('path-group');

    const pathInput = document.createElement('input');
    pathInput.type = 'text';
    pathInput.classList.add(`${type}-path`);
    pathInput.placeholder = `${type.charAt(0).toUpperCase()}${type.slice(1)} path`;
    pathInput.required = true;

    const removeButton = document.createElement('button');
    removeButton.type = 'button';
    removeButton.classList.add('remove-path');
    removeButton.textContent = '-';
    removeButton.addEventListener('click', () => removePath(pathInput));

    pathGroup.appendChild(pathInput);
    pathGroup.appendChild(removeButton);

    const addButton = pathGroup.querySelector('.add-path');
    if (!addButton) {
        const newAddButton = document.createElement('button');
        newAddButton.type = 'button';
        newAddButton.classList.add('add-path');
        newAddButton.textContent = '+';
        newAddButton.addEventListener('click', () => addPath(type));

        pathGroup.appendChild(newAddButton);
    }

    document.getElementById(`${type}-paths`).appendChild(pathGroup);
}

function removePath(pathInput) {
    const pathGroup = pathInput.parentNode;
    pathGroup.removeChild(pathInput);

    const addButtons = pathGroup.querySelectorAll('.add-path');
    if (addButtons.length === 1) {
        addButtons[0].style.display = 'block';
    }
}

document.getElementById('robots-form').addEventListener('submit', (event) => {
    event.preventDefault();

    const allowPaths = Array.from(document.querySelectorAll('.allow-path')).map((pathInput) => pathInput.value.trim());
    const disallowPaths = Array.from(document.querySelectorAll('.disallow-path')).map((pathInput) => pathInput.value.trim());

    const robotsContent = `User-agent: *
Disallow: ${disallowPaths.join('\nDisallow: ')}
Allow: ${allowPaths.join('\nAllow: ')}`;

    document.getElementById('result').value = robotsContent;
});

document.getElementById('generate-button').addEventListener('click', () => {
    const allowPaths = Array.from(document.querySelectorAll('.allow-path')).map((pathInput) => pathInput.value.trim());
    const disallowPaths = Array.from(document.querySelectorAll('.disallow-path')).map((pathInput) => pathInput.value.trim());

    const robotsContent = `User-agent: *
Disallow: ${disallowPaths.join('\nDisallow: ')}
Allow: ${allowPaths.join('\nAllow: ')}`;

    const link = document.createElement('a');
    link.href = `data:text/plain;charset=utf-8,${encodeURIComponent(robotsContent)}`;
    link.download = 'robots.txt';
    link.style.display = 'none';

    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
});
