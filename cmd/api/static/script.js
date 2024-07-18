document.addEventListener('DOMContentLoaded', function() {
    const shortenButton = document.querySelector('#btn-short');
    const linkElement = document.querySelector('#linkText');
    const containerElement = document.querySelector('#shortenedLink');
    const copyButton = document.querySelector('#btnCopy');
    const urlInput = document.querySelector('#urlInput');

    shortenButton.addEventListener('click', function() {
        const url = urlInput.value;

        if (!url) {
            alert('Please enter a URL');
            return;
        }

        fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ original_url: url })
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                linkElement.textContent = data.short_url;
                containerElement.style.display = ''; // Muestra el contenedor
                copyButton.style.display = ''; // Muestra el botón de copiar
            })
            .catch(error => console.error('There has been a problem with your fetch operation:', error));
            fetch('/urlHistory')
    .then(response => response.json())
    .then(data => {
        console.log(Array.isArray(data) && data.length > 0);
        const tableContainer = document.getElementById('urlTableContainer');
        const contentWrapper = document.querySelector('.content-wrapper');
        const mainContent = document.querySelector('.main-content');
        let table = tableContainer.querySelector('table');
        let tbody;

        // Crear la tabla solo si no existe
        if (!table) {
            table = document.createElement('table');
            table.classList.add('table', 'table-bordered');
            const thead = document.createElement('thead');
            tbody = document.createElement('tbody');
            const headerRow = document.createElement('tr');

            headerRow.innerHTML = '<th>URL Original</th><th>URL Corta</th>';
            thead.appendChild(headerRow);
            table.appendChild(thead);
            table.appendChild(tbody);
            tableContainer.appendChild(table);
        } else {
            // Obtener el tbody existente
            tbody = table.querySelector('tbody');
            // Limpiar el tbody existente
            tbody.innerHTML = '';
        }

        if (Array.isArray(data) && data.length > 0) {
            data.forEach(item => {
                const row = document.createElement('tr');
                row.innerHTML = `<td>${item.original_url}</td><td>${item.short_url}</td>`;
                tbody.appendChild(row);
            });
            tableContainer.classList.add('visible');
            mainContent.classList.remove('centered');
        } else {
            console.log('No URL history to display');
            tableContainer.classList.remove('visible');
            mainContent.classList.add('centered');
        }
    })
    .catch(error => console.error('Error fetching URL history:', error));

    });

    copyButton.addEventListener('click', function() {
        // Intenta usar navigator.clipboard primero
        if (navigator.clipboard) {
            navigator.clipboard.writeText(linkElement.textContent)
                .then(() => console.log('Text copied to clipboard'))
                .catch(err => console.error('Failed to copy text: ', err));
        } else {
            // Solución alternativa para navegadores sin soporte para navigator.clipboard
            const textArea = document.createElement('textarea');
            textArea.value = linkElement.textContent;
            document.body.appendChild(textArea);
            textArea.focus();
            textArea.select();
            try {
                const successful = document.execCommand('copy');
                const msg = successful ? 'successful' : 'unsuccessful';
                console.log('Fallback: Copying text command was ' + msg);
            } catch (err) {
                console.error('Fallback: Oops, unable to copy', err);
            }
            document.body.removeChild(textArea);
        }})
});
