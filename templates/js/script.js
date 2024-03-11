function getOrderDetails() {
    // Получаем ключ заказа из текстового поля
    orderKey1 = document.getElementById('orderKey').value;
    document.getElementById("orderKey").onclick = function () { alert('hello!'); };
    alert('Please enter an order key');

    // Проверяем, что ключ заказа был введен
    if (!orderKey1) {
        alert('Please enter an order key');
        return;
    }

    // Строим URL для запроса к серверу
    //const url = /order/${orderKey1};

    // Выполняем асинхронный запрос к серверу
    // fetch(url)
    //     .then(response => {
    //         // Проверяем статус ответа
    //         if (!response.ok) {
    //             throw new Error(Error: ${response.statusText});
    //         }
    //         // Преобразуем ответ в формат JSON
    //         return response.json();
    //     })
    //     .then(data => {
    //         // Формируем HTML с деталями заказа и отображаем его на странице
    //         const details = <p>Order ID: ${data.id}</p>
    //                          <p>Customer Name: ${data.customerName}</p>
    //                          <p>Total Price: ${data.totalPrice}</p>
    //                          <p>Status: ${data.status}</p>;
    //         document.getElementById('orderDetails').innerHTML = details;
    //     })
    //     .catch(error => {
    //         // Обработка ошибок и отображение сообщения об ошибке
    //         console.error('Error fetching order details:', error);
    //         document.getElementById('orderDetails').innerHTML = <p>Error fetching order details</p>;
    //     });
}