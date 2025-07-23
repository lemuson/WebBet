// ВХОД НА СТРАНИЦУ
window.addEventListener("load", () => {
    document.body.classList.add('page-enter');

    const loginForm = document.querySelector('.login-container form');
    const registerForm = document.querySelector('.register-container form');

    loginForm.addEventListener('submit', handleLogin);
    registerForm.addEventListener('submit', handleRegister);
});

// ВХОД
async function handleLogin(event) {
    event.preventDefault();
    const login = document.getElementById("login-username").value;
    const password = document.getElementById("login-password").value;

    try {
        const response = await fetch('/Web-Bet/api/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                login: login,
                password: password
            })
        });

        if (response.ok) {
            window.location.href = '/Web-Bet/profile';
        } else {
            alert("Ошибка входа. Проверьте данные.");
        }
    } catch (error) {
        alert("Произошла ошибка при подключении к серверу.");
    }
}

// РЕГИСТРАЦИЯ
async function handleRegister(event) {
    event.preventDefault();
    const login = document.getElementById("register-username").value;
    const password = document.getElementById("register-password").value;
    const name = document.getElementById("register-name").value;
    const phone = document.getElementById("register-phone").value;

    try {
        const response = await fetch('/Web-Bet/api/users/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                login: login,
                password: password,
                userData: {
                    name: name,
                    phone: phone
                }
            })
        });

        if (response.ok) {
            toggleForm()
        } else {
            alert("Ошибка регистрации. Проверьте данные.");
        }
    } catch (error) {
        alert("Произошла ошибка при подключении к серверу.");
    }
}

// АНИМИРОВАННЫЙ ВЫХОД
function animateExit(event) {
    event.preventDefault();
    document.body.classList.add('page-exit');

    setTimeout(() => {
        window.location.href = event.target.getAttribute('href');
    }, 300);
}

// ВРАЩЕНИЕ КОНТЕЙНЕРА
function toggleForm() {
    const container = document.querySelector('.container');
    container.classList.toggle('rotate');
}
