
document.getElementById("loginForm").addEventListener("submit", async function(event) {
            event.preventDefault();
            const email = document.getElementById("email").value;
            const password = document.getElementById("password").value;

            const data = {
                email: email,
                password: password
            };

            try {
                const response = await fetch("/user_login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(data)
                });
                if (!response.ok) {
                    const errorData = await response.json();
                    throw new Error(errorData.message);
                }

                const responseData = await response.json();
                console.log(responseData); // Output server response
                // Redirect to another page if necessary
                window.location.href = "home.html";
            } catch (error) {
                console.error("There was a problem with your fetch operation:", error);
                alert("Login failed: " + error.message);
            }
        });

