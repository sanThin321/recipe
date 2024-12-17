function logout() {
     var logoutLink = document.getElementById("logoutLink");

    // Add event listener for click event
    logoutLink.addEventListener("click", function(event) {
        event.preventDefault(); // Prevent default link behavior

        fetch('/logout')
            .then(response => {
                if (response.ok) {
                    window.location.href = "/";
                } else {
                    throw new Error(response.statusText);
                }
            })
            .catch(e => {
                alert(e);
            });

        
    });
}