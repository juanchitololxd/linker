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
