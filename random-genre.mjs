import { CountryAdjectives } from './countries.mjs'
import { Nouns } from './nouns.mjs'

const pre = [
    CountryAdjectives,
    [
        'Alternative',
        'Avant-garde',
        'Black',
        'Christian',
        'Crust',
        'Death',
        'Doom',
        'Extreme',
        'Folk',
        'Glam',
        'Gothic',
        'Industrial',
        'Kawaii',
        'Neoclassical',
        'Post',
        'Progressive',
        'Speed',
        'Stoner',
        'Symphonic',
        'Thrash',
    ]
]

const post = [
    ['metal', 'core']
]

function rand(min, max) {
    return Math.floor(Math.random() * (+max - +min)) + +min;
}

function randomFromList(list) {
    return list[rand(0, list.length)]
}

function manyFromLists(lists, min, max) {
    const count = rand(min, max + 1)
    const results = []

    for (let i = 0; i < count; i++) {
        results.push(randomFromList(randomFromList(lists)))
    }
    return results
}

export function randomGenre() {

    return manyFromLists(pre, 1, 3).join(' ') + ' ' + randomFromList(Nouns) + ' ' + manyFromLists(post, 1, 1).join(' ')
}