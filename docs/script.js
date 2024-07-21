document.addEventListener('DOMContentLoaded', () => {
    const sections = document.querySelectorAll('.section');
    sections.forEach(section => section.classList.remove('active'));

    document.getElementById('summary-section').classList.add('active');

   
    document.querySelectorAll('.sidebar a').forEach(link => {
        link.addEventListener('click', event => {
            event.preventDefault(); 
            const sectionId = event.target.getAttribute('data-section');
            sections.forEach(section => section.classList.remove('active')); 
            document.getElementById(`${sectionId}-section`).classList.add('active'); 
        });
    });
});
