import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

with open("requirements.txt", "r") as fh:
    deps = fh.read().split("\n")

setuptools.setup(
    name="thermal-printer-controller", # Replace with your own username
    version="0.0.1",
    author="Thomas Maurice",
    author_email="thomas@maurice.fr",
    description="Thermal printer controller",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/thomas-maurice/thermal-printer",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
    install_requires=deps,
)
