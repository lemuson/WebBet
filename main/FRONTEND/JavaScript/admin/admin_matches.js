//ЗАГРУЗКА СТРАНИЦЫ
document.addEventListener('DOMContentLoaded', () => {
    loadSports();
    loadMatches('');

    document.getElementById('sports-container').addEventListener('click', (event) => {
        const container = event.target.closest('.image-container');
        if (!container) return;

        selectSport(container);
    });
});

//УВЕЛИЧЕНИЕ ВИДА СПОРТА
let expandedElement = null;
function selectSport(container) {
    if (expandedElement === container) {
        container.classList.remove('expanded');
        expandedElement = null;
        loadMatches(``);
    } else {
        if (expandedElement) expandedElement.classList.remove('expanded');
        container.classList.add('expanded');
        expandedElement = container;

        loadMatches(`/sport/${expandedElement.dataset.sport}`);
    }
}

//ЗАГРУЗКА ВИДОВ СПОРТА ИЗ API
async function loadSports() {
    const container = document.getElementById('sports-container');
    container.innerHTML = 'Загрузка...';

    try {
        const res = await fetch('/Web-Bet/api/sports');
        if (!res.ok) throw new Error('Ошибка при загрузке видов спорта');
        const sports = await res.json();

        container.innerHTML = sports.map(sport => `
            <div class="image-container" data-sport="${sport.name}">
                <span class="material-symbols-outlined">${sport.image}</span>
                <span class="name">${sport.name}</span>
            </div>
        `).join('');
    } catch (e) {
        container.textContent = `Не удалось загрузить виды спорта`;
    }
}

//ЗАГРУЗКА МАТЧЕЙ ИЗ API
async function loadMatches(sportPath = '') {
    const container = document.getElementById('matches-container');
    container.innerHTML = 'Загрузка...';

    try {
        const res = await fetch(`/Web-Bet/api/matches${sportPath}`);
        if (!res.ok) throw new Error('Матчи не найдены');
        const matches = await res.json();

        container.innerHTML = matches.map(match => `
            <div class="match-card" data-id="${match.id}">
                <div class="match-info">${match.date}</div>
                <div class="team">
                    <legend>${match.team1.name}</legend>
                    <img src="${match.team1.image}" alt="${match.team1.name}">
                </div>
                <div class="team">
                    <legend>${match.team2.name}</legend>
                    <img src="${match.team2.image}" alt="${match.team2.name}">
                </div>
                <button class="start-button" data-id="${match.id}">НАЧАТЬ</button>
            </div>
        `).join('');

        //ИЗМЕНИТЬ ДАННЫЕ
        document.querySelectorAll('.match-card').forEach(card => {
            card.addEventListener('click', (e) => {
                if (e.target.closest('.start-button')) return;
                const id = card.dataset.id;
                window.location.href = `/Web-Bet/matches/${id}`;
            });
        });

        //"НАЧАТЬ"
        document.querySelectorAll('.start-button').forEach(button => {
            button.addEventListener('click', async (e) => {
                e.stopPropagation();
                const id = button.dataset.id;

                try {
                    const response = await fetch(`/Web-Bet/api/matches/start/${id}`, {
                        method: 'POST'
                    });
                    if (!response.ok) throw new Error('Ошибка запуска матча');
                    alert('Матч запущен');
                } catch (err) {
                    alert('Не удалось запустить матч');
                }
            });
        });

    } catch (e) {
        container.textContent = `Не удалось загрузить матчи`;
    }
}
