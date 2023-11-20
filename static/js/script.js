// let items = document.getElementsByTagName("div")

// for (let i=0; i<items.length; i++) {
//     items[i].addEventListener("click", () => {
//         items[i].classList.toggle("done")
//     })
// }

function toggleEdit(index){
    const taskText =document.getElementById(`taskText${index}`)
    const taskInput =document.querySelector(`#taskForm${index} input[name="updatedTask"]`)
    const updateButton =document.querySelector(`#taskForm${index} button`)

    if(taskText.style.display==="none"){
        taskText.style.display="inline"
        taskInput.style.display="none"
        updateButton.textContent="✏️"
        taskText.textContent=taskInput.value
    }

    else{
        taskText.style.display="none"
        taskInput.style.display="inline"
        updateButton.textContent="✅"
    }
}