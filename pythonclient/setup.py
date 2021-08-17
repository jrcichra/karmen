import setuptools

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setuptools.setup(
    name="karmen",
    version="v2.0.3",
    author="Justin Cichra",
    author_email="justin@none.com",
    description="The karmen python library",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/jrcichra/karmen",
    project_urls={
        "Bug Tracker": "https://github.com/jrcichra/karmen/issues",
    },
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    packages=setuptools.find_packages(where="karmen"),
    python_requires=">=3.7",
)
