
let userEmail = getCookie('email');
window.onload = async function(){
    const data = {
                email: userEmail,
            };

            try {


                const response = await fetch("/get_user", {
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

       
                //profile details
              // Profile details
                const nameElement = document.querySelector('.profile-heading p:nth-child(1) span');
                const emailElement = document.querySelector('.profile-heading p:nth-child(2) span');
                const contactElement = document.querySelector('.profile-heading p:nth-child(3) span');

                // Update the content
                nameElement.textContent = responseData.username;
                emailElement.textContent = responseData.email;
                contactElement.textContent = responseData.number;

                if (responseData.image.length !== 0) {
                    const profileImage = document.getElementById('profileModalBtn');
                    const editprofileImage = document.getElementById('edit_profile_pic');
                    profileImage.src = `data:image/jpeg;base64,${responseData.image}`;
                    editprofileImage.src = `data:image/jpeg;base64,${responseData.image}`;
                }


           
            
            } catch (error) {
                console.error("There was a problem with your fetch operation:", error);
            }
}




function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return null;
}
