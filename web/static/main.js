async function getRequests() {
    const url = "http://localhost:8080/requests";

    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }

        return await response.json;
    } catch(error) {
        console.error(error.message);
    }
}

function populateTable() {
    try {
        const tableBody = document.getElementById("table-body");
        const requests = getRequests();
        const requestData = JSON.parse(requests);

        let rows = "";
        requestData.items.array.forEach(element => {
            const row = document.createElement("tr");

            row.innerHTML = `
                <td>{element.id}</td>
                <td>{element.method}</td>
                <td>{element.path}</td>
                <td>{element.query}</td>
                <td>{element.headers}</td>
                <td>{element.body}</td>
                <td>{element.body_size_bytes}</td>
                <td>{element.body_truncated}</td>
                <td>{element.content_type}</td>
                <td>{element.source_ip}</td>
                <td>{element.received_at}</td>
            `;
            tableBody.appendChild(row);
        });
    } catch(error) {
        tableBody.innerHTML = `<tr><td colspan="3" style="color: red;">{error.errorMessage}</td></tr>`;
    }

}

document.addEventListener("DOMContentLoaded", () => { 
    populateTable();
});