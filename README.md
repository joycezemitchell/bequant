# Bequant

Foobar is a Python library for dealing with word pluralization.

## Installation

Run the following command
```bash
git clone https://github.com/joycezemitchell/bequant.git 
```

## Database Configuration
Open configration.json and update the following

```bash
{
    "host":"xxxx",
	"port":5432,
	"user":"xxxx",
	"password":"xxxx",
	"dbname":"bequant"
}
```


## Usage

```python
import foobar

foobar.pluralize('word') # returns 'words'
foobar.pluralize('goose') # returns 'geese'
foobar.singularize('phenomena') # returns 'phenomenon'
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)