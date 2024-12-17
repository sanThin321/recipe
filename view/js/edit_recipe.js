document.addEventListener('DOMContentLoaded', function () {
    const recipeId = getRecipeIdFromURL();

    if (recipeId) {
        fetchRecipe(recipeId);
    }

    document.getElementById('edit_recipe').addEventListener('submit', function (e) {
        e.preventDefault();
        const formData = new FormData(this);
        const imageInput = document.getElementById('image');
        if (imageInput.files.length === 0) {
            // No new image selected, add the existing image URL to the formData
            formData.append('useExistingImage', true);
        }
        updateRecipe(recipeId, formData);
    });
    document.getElementById('Delete').addEventListener('click', function (e) {
        e.preventDefault();
        deleteRecipe(recipeId);
    });
});

function getRecipeIdFromURL() {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('id');
}

function fetchRecipe(recipeId) {
    fetch(`/recipe/`+ recipeId)
        .then(response => response.json())
        .then(data => {
            showRecipe(data)
        })
        .catch(error => {
            console.error('Error fetching recipe:', error);
        });
}
function showRecipe(data){
    document.getElementById('recipeName').value = data.recipename;
    document.getElementById('ingredients').value = data.ingredient;
    document.getElementById('steps').value = data.steps;
    document.getElementById('image').dataset.currentImage=data.image
    // imageElement.src = data.image;
    // console.log(imageElement.src)
}
function updateRecipe(recipeId, formData) {
    console.log(recipeId)
    fetch(`/recipe/`+ recipeId, {
        method: "PUT",
        body: formData,
    })
    // .then(response => response.json())
    // .then(data => {
    //     console.log('Recipe updated:', data);
    //     alert('Recipe updated successfully!');
    // })
    .then(response => response.text()) // Get the raw response as text
    .then(text => {
        try {
            const data = JSON.parse(text); // Attempt to parse the JSON
            if (data.status === 'success') {
                alert('Recipe updated successfully');
            } else {
                alert('Failed to update recipe');
            }
        } catch (error) {
            console.error('Error parsing response JSON:', error);
            console.log('Response Text:', text); // Log the raw response text
            alert('Failed to update recipe. Server returned invalid JSON.');
        }
    })
    .catch(error => {
        console.error('Error updating recipe:', error);
    });
}
function deleteRecipe(recipeId) {
    if (confirm("Are you sure you want to delete this recipe?")) {
        fetch(`/recipe/`+ recipeId, {
            method: "DELETE",
        })
        .then(response => response.text()) // Get the raw response as text
        .then(text => {
            try {
                const data = JSON.parse(text); // Attempt to parse the JSON
                if (data.status === 'success') {
                    alert('Recipe deleted successfully');
                    window.location.href = '/home.html'; // Redirect to home page after deletion
                } else {
                    alert('Failed to delete recipe');
                }
            } catch (error) {
                console.error('Error parsing response JSON:', error);
                console.log('Response Text:', text); // Log the raw response text
                alert('Failed to delete recipe. Server returned invalid JSON.');
            }
        })
        .catch(error => {
            console.error('Error deleting recipe:', error);
        });
    }
}