const form = document.querySelector('form');
const taskInput = document.querySelector('#task');
const taskList = document.querySelector('ul');
let tasks = [];

function addTask(text) {
    const task = {
        id: Date.now(),
        text: text
    };
    tasks.push(task);
    renderTasks();
    saveTasks();
}

function deleteTask(id) {
    tasks = tasks.filter(task => task.id !== id);
    renderTasks();
    saveTasks();
}

function editTask(id, newText) {
    const index = tasks.findIndex(task => task.id === id);
    if (index !== -1) {
        tasks[index].text = newText;
        renderTasks();
        saveTasks();
    }
}

function renderTasks() {
    taskList.innerHTML = '';
    tasks.forEach(task => {
        const li = document.createElement('li');
        const textSpan = document.createElement('span');
        textSpan.textContent = task.text;
        textSpan.contentEditable = true;
        textSpan.addEventListener('blur', () => {
            editTask(task.id, textSpan.textContent);
        });
        li.appendChild(textSpan);
        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.addEventListener('click', () => {
            deleteTask(task.id);
        });
        li.appendChild(deleteButton);
        taskList.appendChild(li);
    });
}

function saveTasks() {
    localStorage.setItem('tasks', JSON.stringify(tasks));
}

function loadTasks() {
    const storedTasks = localStorage.getItem('tasks');
    if (storedTasks) {
        tasks = JSON.parse(storedTasks);
        renderTasks();
    }
}

form.addEventListener('submit', event => {
    event.preventDefault();
    const text = taskInput.value.trim();
    if (text !== '') {
        addTask(text);
        taskInput.value = '';
        taskInput.focus();
    }
});

loadTasks();
