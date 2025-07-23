document.addEventListener('DOMContentLoaded', () => {
    loadMatchData();
});

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
        }
    } catch (err) {
        console.error("Ошибка запроса:", err);
    }
}