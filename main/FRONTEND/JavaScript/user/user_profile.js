//ЗАГРУЗКА ДАННЫХ
document.addEventListener('DOMContentLoaded', async () => {
    try {
        const res = await fetch('/Web-Bet/api/users/me');
        if (!res.ok) throw new Error('Ошибка при загрузке профиля');
        const user = await res.json();

        //ДАННЫЕ
        document.getElementById('login').value = user.login;
        document.getElementById('name').value = user.userData.name;
        document.getElementById('phone').value = user.userData.phone;
        document.getElementById('balance').textContent = `Баланс: ${user.userData.balance} ₽`;

        //СТАВКИ
        const tbody = document.querySelector('#bets-table tbody');
        tbody.innerHTML = user.bets.map(bet => `
            <tr>
                <td>
                    <a href="/Web-Bet/matches/${bet.matchID}" target="_blank">${bet.match}</a>
                </td>
                <td>${bet.amount} ₽</td>
                <td>${bet.coefficient}</td>
                <td>${bet.prediction}</td>
                <td>${bet.status}</td>
            </tr>
        `).join('');

    } catch (error) {
        console.error('Ошибка загрузки профиля:', error);
        document.querySelector('.profile-container').innerHTML = '<p>Не удалось загрузить данные профиля</p>';
    }
});
