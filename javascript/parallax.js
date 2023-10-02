// parallax.js
window.addEventListener('scroll', function () {
    const scrolled = window.scrollY;
    document.querySelector('.parallax').style.transform = `translateY(-${scrolled * 0.5}px)`; // Ajustez le facteur de vitesse selon vos préférences
});
