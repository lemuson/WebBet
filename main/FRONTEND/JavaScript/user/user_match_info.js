document.addEventListener('DOMContentLoaded', () => {
    loadMatchData();
    loadProfile();
});

//ВЫБОР КОЭФФИЦИЕНТА
function selectOdd(button) {
    document.querySelectorAll('.odds button').forEach(b => b.classList.remove('active'));
    button.classList.add('active');
}

//ФОРМАТ ДАТЫ
function formatDate(dateStr) {
    const date = new Date(dateStr);
    return `${date.toLocaleDateString()} ${date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`;
}

//ЗАГРУЗКА ДАННЫХ О МАТЧЕ ИЗ API
async function loadMatchData() {
    const pathParts = window.location.pathname.split('/');
    const matchId = parseInt(pathParts[pathParts.length - 1]);
    if (isNaN(matchId)) {
        console.error("Некорректный ID матча");
        return;
    }

    try {
        const response = await fetch(`/Web-Bet/api/matches/${matchId}`);
        if (!response.ok) throw new Error('Ошибка при загрузке данных');
        const data = await response.json();
        console.log(data)

        if (!data.team1 || !data.team2 || !Array.isArray(data.predictions)) {
            console.error("Неверный формат данных");
            return;
        }
        document.getElementById('match-date').textContent = `Дата начала матча: ${formatDate(data.date)}`;
        document.getElementById('team1-name').textContent = data.team1.name;
        document.getElementById('team2-name').textContent = data.team2.name;
        document.getElementById('team1-img').src = `${data.team1.image}`;
        document.getElementById('team2-img').src = `${data.team2.image}`;

        const labelsContainer = document.getElementById('labels-container');
        const oddsContainer = document.getElementById('odds-container');
        const matchContainer = document.getElementById('match-container');

        labelsContainer.innerHTML = '';
        oddsContainer.innerHTML = '';

        if (data.result && data.result.trim() !== "") {
            const resultMessage = document.createElement('div');
            resultMessage.className = 'match-result';
            resultMessage.textContent = `МАТЧ ЗАВЕРШЕН. ИСХОД: ${data.result}`;
            oddsContainer.appendChild(resultMessage);
        } else {
            data.predictions.forEach(pred => {
                const label = document.createElement('span');
                label.textContent = pred.name;
                labelsContainer.appendChild(label);

                const btn = document.createElement('button');
                btn.textContent = pred.coefficient;
                btn.setAttribute('onclick', 'selectOdd(this)');
                btn.setAttribute('data-id', pred.id);
                oddsContainer.appendChild(btn);
            });

            matchContainer.innerHTML += `
                <div id="bet-form">
                    <input class="input-amount" type="number" placeholder="Введите сумму ставки">
                    <button onclick="placeBet()">Сделать ставку</button>
                </div>
            `;
        }
    } catch (err) {
        console.error("Ошибка запроса:", err);
    }
}


//ЗАГРУЗКА БАЛАНСА
async function loadProfile(){
    try {
        const res = await fetch('/Web-Bet/api/users/me');
        if (!res.ok) throw new Error('Ошибка получения баланса');
        const user = await res.json();
        document.getElementById('balance').textContent = `Баланс: ${user.userData.balance} ₽`;
    } catch (error) {
        console.error('Ошибка загрузки профиля:', error);
        document.querySelector('.profile-container').innerHTML = '<p>Не удалось загрузить данные профиля</p>';
    }
}

//ОТПРАВКА ДАННЫХ
function placeBet() {
    const selectedBtn = document.querySelector('.odds button.active');
    const amountInput = document.querySelector('.input-amount');
    const amount = parseFloat(amountInput.value);

    if (!selectedBtn || isNaN(amount) || amount <= 0) {
        alert("Выберите коэффициент и введите корректную сумму");
        return;
    }

    const predictionId = parseInt(selectedBtn.getAttribute('data-id'));
    const coefficient = parseFloat(selectedBtn.textContent);

    const betData = {
        prediction_id: predictionId,
        amount: amount,
        coefficient: coefficient
    };

    fetch('/Web-Bet/api/bets', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(betData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Ошибка при отправке ставки");
        }
        return response.json();
    })
    .then(() => {
        alert("Ставка успешно сделана!");
        window.location.href = "/Web-Bet/matches";
    })
    .catch(error => {
        console.error("Ошибка:", error);
        alert("Не удалось сделать ставку.");
    });
}