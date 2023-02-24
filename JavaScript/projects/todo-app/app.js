const addButton = document.getElementById('addButton');
const taskInput = document.getElementById('taskInput');
const taskList = document.getElementById('taskList');

addButton.addEventListener('click', function() {
    const task = document.createElement('li');
    task.textContent = taskInput.value;
    taskList.appendChild(task);
    taskInput.value = '';
});

taskList.addEventListener('click', function(event) {
    if (event.target.tagName === 'LI') {
        event.target.parentNode.removeChild(event.target);
    }
});

// Save tasks to local storage
function saveTasks() {
    localStorage.setItem('tasks', taskList.innerHTML);
}

// Load tasks from local storage
function loadTasks() {
    taskList.innerHTML = localStorage.getItem('tasks');
}

// Add event listener to save tasks
window.addEventListener('beforeunload', saveTasks);

// Load saved tasks
window.addEventListener('load', loadTasks);