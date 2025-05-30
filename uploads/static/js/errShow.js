const errorDiv = document.getElementById('error-message');
if (errorDiv) {
    setTimeout(() => {
        errorDiv.style.opacity = '0';
        setTimeout(() => errorDiv.remove(), 300);
    }, 3000);
}