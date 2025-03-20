// cursor--Initial commit: Add JavaScript to handle form submission and API calls.
// cursor--Update: Change API URL to absolute endpoint and use modal for feedback.
// cursor--SweetAlert2 integration: Replace custom modal feedback with SweetAlert2 popups.
// cursor--SweetAlert2 theme integration: Create a mixin to enforce theme styling.
const themedSwal = Swal.mixin({
  customClass: {
    popup: 'altheme-swal-popup',
    confirmButton: 'altheme-swal-confirm-button'
  },
  buttonsStyling: false,
  background: '#e7d9c5', // Warm Beige
  color: '#3e2f1e'       // Dark Brown
});

// cursor--Add DOMContentLoaded wrapper to ensure the DOM is loaded before attaching event listeners.
document.addEventListener('DOMContentLoaded', function() {
  const signupForm = document.getElementById('signup-form');
  if (!signupForm) {
    console.error('signup-form element not found');
    return;
  }
  
  signupForm.addEventListener('submit', async function(e) {
    e.preventDefault();
    const email = document.getElementById('email').value;
  
    try {
      // Use an absolute URL so that the request goes to the backend container on port 8080.
      const response = await fetch('http://localhost:8080/subscribe', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: email })
      });
  
      const data = await response.json();
  
      if (response.ok) {
        themedSwal.fire({ title: 'Success!', text: data.message, icon: 'success', confirmButtonText: 'OK' });
      } else {
        themedSwal.fire({ title: 'Error', text: data.message || 'An error occurred', icon: 'error', confirmButtonText: 'Try Again' });
      }
      
      // Clear the email input
      document.getElementById('email').value = '';
    } catch (error) {
      themedSwal.fire({ title: 'Network Error', text: 'Network error. Please try again later.', icon: 'error', confirmButtonText: 'OK' });
    }
  });
});

// cursor--Remove obsolete modal event listener since modal container was removed.
// Added event listener to close modal when clicking outside the modal content.
// window.addEventListener('click', function(event) {
//   const modal = document.getElementById('modal');
//   if (event.target === modal) {
//     modal.style.display = 'none';
//   }
// }); 