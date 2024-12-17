document.getElementById("registration").addEventListener("submit", async function(event) {
    event.preventDefault(); // Prevent the form from submitting

    // Get input values directly
    const username = document.getElementById("username").value;
    const email = document.getElementById("email").value;
    const number = document.getElementById("number").value;
    const password = document.getElementById("password").value;

    const userData = {
        username: username,
        email: email,
        number: number,
        password: password
    };

    try {
        const response = await fetch("/user_register", {
            method: "POST",
            body: JSON.stringify(userData),
            headers: {
                "Content-Type": "application/json"
            }
        });

        if (!response.ok) {
            const err = await response.json();
            throw new Error(err.message);
        }

       
        alert("Registration successful!, You can now Login");

        window.location.href = "/";

    } catch (error) {
        console.error("There was a problem with your fetch operation:", error);
        alert("Registration failed: " + error.message);
    }
});
