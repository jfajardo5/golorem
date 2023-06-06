<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="left">

<h3 align="left">Golorem</h3>

  <p align="left">
    Golorem is a super fast and convenient lorem ipsum CLI generator.
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Golorem

Generate lorem ipsum with a quick terminal command

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [Dolorem](https://github.com/jfajardo5/dolorem)
* [Go](https://go.dev/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

<p>Grab the latest build from the 'build' directory and place it in your PATH.</p>

~~~
➜  ~ golorem
Dolorem ipsum dolor sit amet, insideo multo urbs infirmo advenio innuo dignitas, demens infamo sapienter procul coepi infirmatio inceptor,  devoveo proinde sapientia suffragium sequor aliquantus, spectaculum informis curiositas aestivus pollen penitus deduco, postea facio spiculum noceo aegrus novitas uberrime, quisquam principatus usus deputo orno levitas supero, apparatus quantuvis illos insurgo redigo emendo insciens.
~~~



### Installation

  * You can just grab the executable from this repo's 'build' directory and place it in your PATH.
  * You may also clone this repo and build it yourself by running the following command in the project folder:
~~~
$ go build -o build/golorem .
~~~

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

~~~
➜  ~ golorem -h
Usage of golorem:
  -help
    	View usage
  -p	Generate paragraphs. Add extra numeric parameters to override default:
    	* X (Number of paragraphs)
    	* Y (Number of sentences per paragraph)
    	* Z (Number of words per sentence)
  -s	Generate a sentence. Add an extra numeric parameter to override default:
    	* X (Number of words in the sentece)
  -w	Generate a word.
~~~


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>







<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/jfajardo5/golorem.svg?style=for-the-badge
[contributors-url]: https://github.com/jfajardo5/golorem/graphs/contributors
[stars-shield]: https://img.shields.io/github/stars/jfajardo5/golorem.svg?style=for-the-badge
[stars-url]: https://github.com/jfajardo5/golorem/stargazers
[issues-shield]: https://img.shields.io/github/issues/jfajardo5/golorem.svg?style=for-the-badge
[issues-url]: https://github.com/jfajardo5/golorem/issues
[license-shield]: https://img.shields.io/github/license/jfajardo5/golorem.svg?style=for-the-badge
[license-url]: https://github.com/jfajardo5/golorem/blob/master/LICENSE.txt
