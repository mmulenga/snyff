const tableBody = document.getElementById("table-body");

async function getRequests() {
    const url = "http://localhost:8080/requests";

    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
    }

    let r = await response.json();

    return r;
}

async function populateTable() {
    try {
        const requests = await getRequests();
        const decoder = new TextDecoder('utf-8')
        let rows = "";

        requests.forEach(request => {
            const row = document.createElement("tr");
            Object.keys(request).forEach(key => {
                const newCell = row.insertCell();

                switch (key) {
                    case 'headers':
                        Object.entries(request[key]).forEach(header => {
                            newCell.textContent += header + '\n';
                        })
                        break;
                    case 'body':
                        newCell.textContent = atob(request[key]);
                        break;
                    default:
                        newCell.textContent = request[key];
                        break;
                }
            })
            tableBody.appendChild(row);
        });
    } catch(error) {
        tableBody.textContent = `<tr><td colspan="3" style="color: red;">${error.message}</td></tr>`;
    }

}

populateTable();