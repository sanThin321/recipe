
  document.getElementById('uploadForm').addEventListener('submit', function(e) {
    e.preventDefault();
    
    let user_id = getCookie('id')
    
    const file =  document.getElementById('profile_image').files[0];
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const contact = document.getElementById('contact').value; 
    
    console.log(file)
    
    if (file) {
        // console.log("here")
            const reader = new FileReader();
            reader.onloadend = () => {
                const base64Profile = reader.result;
                const profileObj = {
                    profilepicture: base64Profile,
                    name: name,
                    email:email,
                    contact: contact,
                };

                console.log("here")
                
                fetch('/update_user/' + user_id, {
                    method: "PUT",
                    body: JSON.stringify(profileObj),
                    headers: {
                        "Content-Type": "application/json; charset=UTF-8"
                    },
                })
                .then((res) => {
                    if (res.ok) {
                      alert("Please login again")
                      window.location.href = '/';
                    } else {
                        alert("Server: update request error");
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
            };
            reader.readAsDataURL(file);
        }else{
          console.log("no file")
        }
});

function deleteItem(){
    let user_id = getCookie('id')
    if (confirm("Are you sure you want to delete this item?")) {
                // Code to delete the item
                  fetch('/del_user/' + user_id, {
                    method: "DELETE",
                })
                .then((res) => {
                    if (res.ok) {
                      alert("User deleted!");
                      alert("Please login again")
                      window.location.href = '/';
                    } else {
                        alert("Server: update request error");
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
            } else {
                alert("Deletion canceled.");
            }

   
}