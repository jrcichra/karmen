from distutils.core import setup
setup(
    name='karmen',         # How you named your package folder (MyLib)
    packages=['karmen'],   # Chose the same as "name"
    version='0.2',      # Start with a small number and increase it with every change you make
    # Chose a license from here: https://help.github.com/articles/licensing-a-repository
    license='MIT',
    # Give a short description about your library
    description='A centralized microservices framework for declarative applications',
    author='Justin Cichra',                   # Type in your name
    author_email='jrcichra@yahoo.com',      # Type in your E-Mail
    # Provide either the link to your github or to your website
    url='https://github.com/jrcichra/karmen',
    # I explain this later on
    download_url='https://github.com/jrcichra/karmen/archive/v_01.tar.gz',
    # Keywords that define your package best
    keywords=['Docker', 'Message', 'Queue'],
    install_requires=[
    ],
    classifiers=[
        # Chose either "3 - Alpha", "4 - Beta" or "5 - Production/Stable" as the current state of your package
        'Development Status :: 4 - Beta',
        # Define that your audience are developers
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Build Tools',
        'License :: OSI Approved :: MIT License',   # Again, pick a license
        # Specify which pyhton versions that you want to support
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
    ],
)
