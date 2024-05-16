document.addEventListener('DOMContentLoaded', function() {
    const viewMoreBtn = document.querySelector('.view-more-btn');
    const additionalDetails = document.querySelector('.additional-details');
  
    viewMoreBtn.addEventListener('click', function() {
      additionalDetails.classList.toggle('show');
      if (additionalDetails.classList.contains('show')) {
        viewMoreBtn.textContent = 'View Less';
      } else {
        viewMoreBtn.textContent = 'View More';
      }
    });
  });
  