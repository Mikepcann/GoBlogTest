import axios from 'axios'

const SERVER = 'http://localhost:8080/'

export const get = async (url) => {
	const fullUrl = `${SERVER}${url}`
	try {
		console.log("the URL",fullUrl)
		const result = await axios.get(fullUrl)
		return result
	} catch (err) {
		console.log('API GET CAll ERROR: ', err.message)
		return err
	}
}

export const post = async (url, val) => {
	const fullUrl = `${SERVER}${url}`
    try {
        const result = await axios.post(fullUrl, val)
        return result
    } catch (err) {
        console.log('API POST CAll ERROR: ', err.message)
		return err
    }
}

export const deleteMethod = async () => {
	const fullUrl = `${SERVER}posts`
	try {
		const result = await axios.delete(fullUrl)
		return result
	} catch (err) {
		console.log('API GET CAll ERROR: ', err.message)
		return err
	}
}
export const deleteById = async (id) => {
	const fullUrl = `${SERVER}posts/${id}`
	try {
		const result = await axios.delete(fullUrl)
		return result
	} catch (err) {
		console.log('API GET CAll ERROR: ', err.message)
		return err
	}
}
