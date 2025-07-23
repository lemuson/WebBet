document.addEventListener("DOMContentLoaded", () => {
    loadSports();
    loadTeams();
    loadResults();

    document.getElementById("create-sport-form").addEventListener("submit", handleCreateSport);
    document.getElementById("create-team-form").addEventListener("submit", handleCreateTeam);
    document.getElementById("create-match-form").addEventListener("submit", handleCreateMatch);
});

// /Web-Bet/api/sports
async function loadSports() {
    try {
        const response = await fetch('/Web-Bet/api/sports');
        const sports = await response.json();
        const sportSelect = document.getElementById('sport-select');
        sports.forEach(sport => {
            const option = document.createElement('option');
            option.value = sport.id;
            option.textContent = sport.name;
            sportSelect.appendChild(option);
        });
    } catch (error) {
        console.error('Ошибка при загрузке видов спорта:', error);
    }
}

// /Web-Bet/api/teams
async function loadTeams() {
    try {
        const response = await fetch('/Web-Bet/api/teams');
        const teams = await response.json();
        const team1Select = document.getElementById('team1-select');
        const team2Select = document.getElementById('team2-select');
        teams.forEach(team => {
            const option = document.createElement('option');
            option.value = team.id;
            option.textContent = team.name;
            team1Select.appendChild(option);
            team2Select.appendChild(option.cloneNode(true));
        });
    } catch (error) {
        console.error('Ошибка при загрузке команд:', error);
    }
}

// /Web-Bet/api/results
async function loadResults() {
    try {
        const response = await fetch('/Web-Bet/api/results');
        const results = await response.json();
        const oddsContainer = document.getElementById('odds');
        results.forEach(result => {
            const div = document.createElement('div');
            div.classList.add('odd');
            div.innerHTML = `
                <label for="result-${result.id}">${result.name}</label>
                <input type="number" id="result-${result.id}" step="0.01" inputmode="decimal">
            `;
            oddsContainer.appendChild(div);
        });
    } catch (error) {
        console.error('Ошибка при загрузке результатов:', error);
    }
}

// === ОБРАБОТЧИКИ ФОРМ ===
async function handleCreateSport(event) {
    event.preventDefault();
    const form = event.target;
    const name = document.getElementById("sport-name").value;
    const image = document.getElementById("sport-image").value;

    const success = await postData("/Web-Bet/api/sports", { name, image });
    if (success) {
        alert("Вид спорта успешно создан!");
        form.reset();
    }
}

async function handleCreateTeam(event) {
    event.preventDefault();

    const name = document.getElementById("team-name").value;
    const fileInput = document.getElementById("team-image");
    const file = fileInput.files[0];

    const formData = new FormData();
    formData.append("image", file);

    // 1. Отправляем файл на сервер
    const uploadRes = await fetch("/Web-Bet/upload-image", {
        method: "POST",
        body: formData
    });

    const uploadData = await uploadRes.json();

    if (!uploadRes.ok) {
        alert("Ошибка загрузки изображения: " + uploadData.error);
        return;
    }

    const teamRes = await fetch("/Web-Bet/api/teams", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            name: name,
            image: uploadData.path
        })
    });

    const teamData = await teamRes.json();

    if (!teamRes.ok) {
        alert("Ошибка сохранения команды: " + teamData.error);
    } else {
        alert("Команда успешно сохранена!");
    }
}

async function handleCreateMatch(event) {
    event.preventDefault();
    const form = event.target;
    const date = document.getElementById("match-date").value;
    const sport_id = parseInt(document.getElementById("sport-select").value);
    const team1_id = parseInt(document.getElementById("team1-select").value);
    const team2_id = parseInt(document.getElementById("team2-select").value);
    const predictions = getPredictionsFromOdds();
    const matchData = {
        date,
        team1_id,
        team2_id,
        sport_id,
        result_id: null,
        predictions
    };

    console.log(matchData);
    const success = await postData("/Web-Bet/api/matches", matchData);
    if (success) {
        alert("Матч успешно создан!");
        form.reset();
        clearOddsFields();
    }
}


async function postData(url, data) {
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            const errorText = await response.text();
            console.error(`Ошибка при POST-запросе: ${errorText}`);
            alert("Ошибка при сохранении. Проверьте введённые данные.");
            return false;
        } else {
            return true;
        }
    } catch (error) {
        console.error("Ошибка при отправке запроса:", error);
        alert("Ошибка соединения с сервером.");
        return false;
    }
}

function getPredictionsFromOdds() {
    const oddsElements = document.querySelectorAll('.odd input[type="number"]');
    return Array.from(oddsElements).map(input => ({
        result_id: parseInt(input.id.replace('result-', '')),
        coefficient: parseFloat(input.value) || 0
    }));
}

function clearOddsFields() {
    const oddsElements = document.querySelectorAll("#odds .input-group");
    oddsElements.forEach(el => {
        el.querySelector(".result-id").value = "";
        el.querySelector(".coefficient").value = "";
    });
}
