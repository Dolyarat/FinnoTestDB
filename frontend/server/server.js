import express from 'express';
import axios from 'axios';
import cors from 'cors';
import bodyParser from 'body-parser';
const app = express();

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended:true}));

app.get('/persons', async (req, res) => {
    try {
        const apiData = await axios.get(`http://localhost:5000/persons`);
        res.json(apiData.data);
    } catch (error) {
        console.error('Error:', error);
        res.status(500).send('Internal Server Error');
    }
});

app.get('/persons/:id', async (req, res) => {
    try {
        const id = req.params.id
        const apiData = await axios.get(`http://localhost:5000/persons/${id}`);
        res.json(apiData.data);
    } catch (error) {
        console.error('Error:', error);
        res.status(500).send('Internal Server Error');
    }
});

app.post('/persons', async (req, res) => {
    try {
        const postData = await axios.post(`http://localhost:5000/persons`, req.body);
        res.status(200).send('Success Post');
    } catch (error) {
        console.error('Error:', error);
        res.status(500).send('Internal Server Error');
    }
});

app.put('/persons/:id', async (req, res) => {
    try {
        const id = req.params.id
        const putData = await axios.put(`http://localhost:5000/persons/${id}`, req.body);
        res.status(200).send('Success Put');
    } catch (error) {
        console.error('Error:', error);
        res.status(500).send('Internal Server Error');
    }
});

app.delete('/persons/:id', async (req, res) => {
    try {
        const id = req.params.id
        const delData = await axios.delete(`http://localhost:5000/persons/${id}`);
        res.status(200).send('Success Delete');
    } catch (error) {
        console.error('Error:', error);
        res.status(500).send('Internal Server Error');
    }
});

app.listen(8888, () => {
    console.log('Server start')
});