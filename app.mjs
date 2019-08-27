import { randomGenre } from './random-genre.mjs'

const genre = document.getElementById('genre')
const regenerate = document.getElementById('regenerate')

genre.innerText = randomGenre()

regenerate.onclick = () => {
    genre.innerText = randomGenre()
}
